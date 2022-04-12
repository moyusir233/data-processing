package test

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/conf"
	"gitee.com/moyusir/data-processing/internal/data"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/golang/protobuf/proto"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		"wFt-QqAqn_q5gWnCkPtJbewTUnkHhUDKSo0XUIUZtJliMaf6vnrneaxymtwvT34Yzu2xi2ClDXlTWLF9gb-otQ=="

	redisData := &data.RedisData{}

	influxdbData, cleanUp2, err := data.NewInfluxdbData(bootstrap.Data)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		defer cleanUp2()
		client := influxdb2.NewClient(
			bootstrap.Data.Influxdb.ServerUrl, bootstrap.Data.Influxdb.AuthToken)
		defer client.Close()

		deleteAPI := client.DeleteAPI()
		start := time.Now().UTC().Add(-1 * time.Hour)
		end := time.Now().UTC().Add(3 * time.Hour)
		for id := 0; id < 2; id++ {
			predicate := fmt.Sprintf(`deviceClassID="%d"`, id)
			for _, bucket := range []string{"test", "test-warning_detect", "test-warnings"} {
				deleteAPI.DeleteWithName(context.Background(), "test", bucket, start, end, predicate)
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
		err = repo.SaveWarningMessage("test", warnings...)
		if err != nil {
			t.Fatal(err)
		}

		// 测试分页查询
		stop := now.Add(10 * time.Minute)
		option := &biz.QueryOption{
			Bucket:     "test",
			Start:      &now,
			Stop:       &stop,
			CountQuery: false,
			// 注意limit和offset都要依据warning的field的数量放缩
			Limit:      3,
			Offset:     0,
			GroupCount: 3,
		}

		for i, warning := range warnings {
			option.Offset = i * 3
			messages, err := repo.GetWarningMessage(option)
			if err != nil {
				t.Fatal(err)
			}

			if len(messages) != 1 {
				t.Error("分页查询的limit失效了")
				continue
			}

			// 由于timestamppb保存的时间戳是当做UTC时间保存的，如果希望protobuf的timestamp对象
			// 解析出来的时间日期时本地时间，则需要对时间进行偏移
			warning.Start = timestamppb.New(warning.Start.AsTime().Add(8 * time.Hour))
			warning.End = timestamppb.New(warning.End.AsTime().Add(8 * time.Hour))

			// 序列化为json后再进行比较
			warningJson, err := protojson.Marshal(warning)
			if err != nil {
				t.Fatal(err)
			}
			msgJson, err := protojson.Marshal(messages[0])
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(warningJson, msgJson) {
				t.Errorf(
					"分页查询得到的警告信息与期望的不一致:\nexpect:\n%v\npractice:\n%v",
					string(warningJson), string(msgJson),
				)
			}
			//if !proto.Equal(warning, messages[0]) {
			//	t.Errorf(
			//		"分页查询得到的警告信息与期望的不一致:\nexpect:\n%v\npractice:\n%v",
			//		warning, messages[0],
			//	)
			//	marshal, err := protojson.Marshal(messages[0])
			//	if err != nil {
			//		continue
			//	}
			//	t.Error(string(marshal))
			//}
		}

		// 测试查询记录总数
		count, err := repo.GetRecordCount(option)
		if err != nil {
			t.Fatal(err)
		}

		if int(count/3) != len(warnings) {
			t.Fatalf("查询到的记录数量与期望的不符合:%d,%d", count/3, len(warnings))
		} else {
			t.Log("查询到的记录数量正确")
		}

	})

}
