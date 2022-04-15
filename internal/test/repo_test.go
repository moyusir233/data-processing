package test

import (
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/conf"
	"gitee.com/moyusir/data-processing/internal/data"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/golang/protobuf/proto"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"testing"
	"time"
)

func TestRepo_Influxdb(t *testing.T) {
	bootstrap, err := conf.LoadConfig("../../configs/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	bootstrap.Data.Influxdb.ServerUrl = "http://localhost:8086"
	bootstrap.Data.Influxdb.AuthToken =
		"33124V3gfLPi2wy07KIJvPSQdIey9ogLYh6AHhtjwKJgyg1Xguy-ApjUh4VWcZUw6jCCfinxTla_uZdHxIWaEw=="
	initBucket := "test"

	redisData := &data.RedisData{}

	// 用于写入数据和删除数据的influxdb客户端
	client := influxdb2.NewClient(
		bootstrap.Data.Influxdb.ServerUrl, bootstrap.Data.Influxdb.AuthToken)
	t.Cleanup(func() {
		client.Close()
	})

	influxdbData, cleanUp2, err := data.NewInfluxdbData(bootstrap.Data)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		defer cleanUp2()

		deleteAPI := client.DeleteAPI()
		start := time.Now().UTC().Add(-1 * time.Hour)
		end := time.Now().UTC().Add(3 * time.Hour)
		for id := 0; id < 2; id++ {
			predicate := fmt.Sprintf(`deviceClassID="%d"`, id)
			for _, bucket := range []string{initBucket, initBucket + "-warning_detect", initBucket + "-warnings"} {
				deleteAPI.DeleteWithName(
					context.Background(), bootstrap.Data.Influxdb.Org, bucket, start, end, predicate)
			}
		}
	})

	repo := data.NewRepo(redisData, influxdbData)

	t.Run("Test_PageQuery", func(t *testing.T) {
		// 用于创建结构体副本的辅助函数
		copyAndAssign := func(src proto.Message, fieldName string, fieldValue interface{}) proto.Message {
			dst := proto.Clone(src)
			messageReflect := proto.MessageReflect(dst)
			fieldDescriptor := messageReflect.Descriptor().Fields().ByTextName(fieldName)

			if message, ok := fieldValue.(proto.Message); ok {
				messageReflect.Set(
					fieldDescriptor, protoreflect.ValueOfMessage(proto.MessageReflect(message)))
			} else {
				messageReflect.Set(fieldDescriptor, protoreflect.ValueOf(fieldValue))
			}
			return dst
		}

		// 准备用于分页测试的警告信息数据
		now := time.Now()
		warnings := []*utilApi.Warning{
			{
				DeviceClassId:   0,
				DeviceId:        "test1",
				DeviceFieldName: "current",
				WarningMessage:  "warning",
				Start:           timestamppb.New(now),
				End:             timestamppb.New(now.Add(5 * time.Minute)),
				Processed:       false,
			},
		}

		warnings = append(warnings,
			copyAndAssign(
				warnings[0], "start", timestamppb.New(now.Add(5*time.Minute))).(*utilApi.Warning),
			copyAndAssign(warnings[0], "device_field_name", "voltage").(*utilApi.Warning),
			copyAndAssign(warnings[0], "device_id", "test2").(*utilApi.Warning),
			copyAndAssign(warnings[0], "device_class_id", int32(1)).(*utilApi.Warning),
		)

		// 写入测试数据
		err = repo.SaveWarningMessage(initBucket, warnings...)
		if err != nil {
			t.Fatal(err)
		}

		// 等待写入
		time.Sleep(time.Second)
		// 测试分页查询
		stop := now.Add(10 * time.Minute)
		option := &biz.QueryOption{
			Bucket:     initBucket,
			Start:      &now,
			Stop:       &stop,
			CountQuery: false,
			Filter: map[string]string{
				"processed": strconv.FormatBool(false),
			},
			// 注意limit和offset都要依据warning的field的数量放缩
			Limit:      3,
			Offset:     0,
			GroupCount: 3,
		}

		for i, warning := range warnings {
			option.Offset = i * 3

			messages, err := repo.GetWarningMessage(*option)
			if err != nil {
				t.Fatal(err)
			}

			if len(messages) != 1 {
				t.Error("分页查询的limit失效了")
				continue
			}

			equal := true
			switch msg := messages[0]; {
			case warning.DeviceId != msg.DeviceId:
				t.Error("设备id不同")
				equal = false
			case warning.DeviceClassId != msg.DeviceClassId:
				t.Error("设备类别id不同")
				equal = false
			case warning.Start.AsTime().Format(time.RFC3339) != msg.Start.AsTime().Format(time.RFC3339):
				t.Error("警告起始时间不同")
				equal = false
			case warning.End.AsTime().Format(time.RFC3339) != msg.End.AsTime().Format(time.RFC3339):
				t.Error("警告结束时间不同")
				equal = false
			case warning.DeviceFieldName != msg.Tags["deviceFieldName"]:
				t.Error("警告字段不同")
				equal = false
			case warning.WarningMessage != msg.Fields["warningMessage"]:
				t.Error("警告信息不同")
				equal = false
			}
			if !equal {
				t.Errorf(
					"期望的警告信息与实际查询得到的不符:\nexpect:\n%v\npractice:\n%v",
					warning, messages[0],
				)
			}
		}

		// 测试查询记录总数
		count, err := repo.GetRecordCount(*option)
		if err != nil {
			t.Fatal(err)
		}

		if int(count/3) != len(warnings) {
			t.Fatalf("查询到的记录数量与期望的不符合:%d,%d", count/3, len(warnings))
		} else {
			t.Log("查询到的记录数量正确")
		}

		t.Run("Test_UpdateAndRemoveWarning", func(t *testing.T) {
			// 先更新设备处理状态
			for _, w := range warnings {
				err := repo.UpdateWarningProcessedState(initBucket, &v1.UpdateWarningRequest{
					DeviceClassId:   w.DeviceClassId,
					DeviceId:        w.DeviceId,
					DeviceFieldName: w.DeviceFieldName,
					Time:            w.Start,
					Processed:       true,
				})
				if err != nil {
					t.Fatal(err)
				}
			}
			// 查询，检验更新是否成功
			option.Limit = 0
			option.Filter["processed"] = strconv.FormatBool(true)
			replyWarnings, err := repo.GetWarningMessage(*option)
			if err != nil {
				t.Fatal(err)
			}
			if len(replyWarnings) != len(warnings) {
				t.Fatal("查询得到的警告信息数量错误")
			}
			for _, w := range replyWarnings {
				if w.Tags["processed"] == strconv.FormatBool(false) {
					t.Error("警告信息处理状态更新失败")
				}
			}

			// 删除掉所有警告信息
			for _, w := range warnings {
				err := repo.DeleteWarningMessage(initBucket, &v1.DeleteWarningRequest{
					DeviceClassId:   w.DeviceClassId,
					DeviceId:        w.DeviceId,
					DeviceFieldName: w.DeviceFieldName,
					Time:            w.Start,
				})
				if err != nil {
					t.Fatal(err)
				}
			}

			// 查询，检验是否删除成功
			replyWarnings, err = repo.GetWarningMessage(*option)
			if err != nil {
				t.Fatal(err)
			}
			if len(replyWarnings) != 0 {
				t.Error("警告信息删除失败:")
				for _, w := range replyWarnings {
					marshal, err := protojson.Marshal(w)
					if err != nil {
						continue
					}
					t.Logf("%s\n", string(marshal))
				}
				t.Fatal()
			}
		})
	})

	t.Run("Test_RemoveDeviceState", func(t *testing.T) {
		// 准备写入的设备状态数据
		now := time.Now().UTC()
		state := &v1.DeviceState{
			DeviceClassId: 1,
			DeviceId:      "test",
			Time:          timestamppb.New(now),
			Fields: map[string]float64{
				"current": 200,
			},
			Tags: map[string]string{
				"deviceType": "computer",
			},
		}
		// 写入设备状态
		writeAPI := client.WriteAPI(bootstrap.Data.Influxdb.Org, initBucket)
		point := write.NewPointWithMeasurement(state.DeviceId).
			AddTag("deviceClassID", strconv.Itoa(int(state.DeviceClassId))).
			SetTime(state.Time.AsTime())
		for k, v := range state.Tags {
			point.AddTag(k, v)
		}
		for k, v := range state.Fields {
			point.AddField(k, v)
		}
		writeAPI.WritePoint(point)
		writeAPI.Flush()

		// 查询，检测写入是否成功
		// 等待写入
		time.Sleep(time.Second)
		option := &biz.QueryOption{
			Bucket:     initBucket,
			Past:       time.Hour,
			GroupCount: len(state.Fields),
			Filter: map[string]string{
				"deviceType": "computer",
			},
		}
		info, err := repo.BatchGetDeviceStateInfo(int(state.DeviceClassId), *option)
		if err != nil {
			t.Fatal(err)
		}
		if len(info) != 1 {
			t.Fatalf("查询到的设备状态数量错误:%d", len(info))
		}
		t.Log(info[0])

		// 删除设备状态
		err = repo.DeleteDeviceStateInfo(initBucket, &v1.DeleteDeviceStateRequest{
			DeviceClassId: state.DeviceClassId,
			DeviceId:      state.DeviceId,
			Time:          state.Time,
		})
		if err != nil {
			t.Fatal(err)
		}

		// 再次查询，检验是否删除成功
		info, err = repo.BatchGetDeviceStateInfo(int(state.DeviceClassId), *option)
		if err != nil {
			t.Fatal(err)
		}
		if len(info) != 0 {
			t.Fatal("删除设备状态失败")
		}
	})
}
