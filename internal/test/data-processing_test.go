package test

import (
	"bytes"
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/data"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"testing"
	"time"
)

func TestDataProcessingService(t *testing.T) {
	// 进行的测试包括:
	// 1. 测试指定设备的配置查询功能
	// 2. 测试设备状态的批量查询功能
	// 3. 测试关于设备状态的注册信息的查询功能
	// 4. 测试预警检测与多连接的预警推送功能
	// 5. 测试警告信息的批量查询功能

	// 定义测试时使用的注册信息
	registerInfo := []utilApi.DeviceStateRegisterInfo{
		{
			// 以utilApi.TestedDeviceState为模板进行定义
			Fields: []*utilApi.DeviceStateRegisterInfo_Field{
				// 定义字段voltage的预警规则为1秒内的平均值不超过100
				{
					Name: "voltage",
					Type: utilApi.Type_DOUBLE,
					WarningRule: &utilApi.DeviceStateRegisterInfo_WarningRule{
						CmpRule: &utilApi.DeviceStateRegisterInfo_CmpRule{
							Cmp: utilApi.DeviceStateRegisterInfo_GT,
							Arg: "100",
						},
						AggregationOperation: utilApi.DeviceStateRegisterInfo_AVG,
						Duration:             durationpb.New(time.Second),
					},
				},
				// 定义字段current的预警规则为1秒内的最大值不小于100
				{
					Name: "current",
					Type: utilApi.Type_DOUBLE,
					WarningRule: &utilApi.DeviceStateRegisterInfo_WarningRule{
						CmpRule: &utilApi.DeviceStateRegisterInfo_CmpRule{
							Cmp: utilApi.DeviceStateRegisterInfo_LT,
							Arg: "100",
						},
						AggregationOperation: utilApi.DeviceStateRegisterInfo_MAX,
						Duration:             durationpb.New(time.Second),
					},
				},
				// 定义字段temperature的预警规则为1秒内的最小值不等于100
				{
					Name: "temperature",
					Type: utilApi.Type_DOUBLE,
					WarningRule: &utilApi.DeviceStateRegisterInfo_WarningRule{
						CmpRule: &utilApi.DeviceStateRegisterInfo_CmpRule{
							Cmp: utilApi.DeviceStateRegisterInfo_EQ,
							Arg: "100",
						},
						AggregationOperation: utilApi.DeviceStateRegisterInfo_MIN,
						Duration:             durationpb.New(time.Second),
					},
				},
			},
		},
	}

	// 初始化测试环境
	bootstrap, err := generalInit("")
	if err != nil {
		t.Fatal(err)
	}

	// 初始化用于写入测试用例信息的redis客户端
	redisClient, cleanup, err := data.NewData(bootstrap.Data, log.NewStdLogger(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		// 测试结束时清空数据库
		redisClient.FlushDB(context.Background())
		cleanup()
	})

	// 启动服务器，并实例化测试所需的http客户端
	configHTTPClient, warningDetectHTTPClient := StartDataProcessingServer(t, bootstrap, registerInfo)

	// 1. 测试设备的配置查询功能
	t.Run("Test_GetDeviceConfig", func(t *testing.T) {
		// 初始化测试数据，并写入到数据库中
		var configs []utilApi.TestedDeviceConfig
		for i := 0; i < 5; i++ {
			config := utilApi.TestedDeviceConfig{
				Id:     fmt.Sprintf("test%d", i),
				Status: false,
			}
			err := saveConfig(redisClient, config.Id, &config)
			if err != nil {
				t.Error(err)
				return
			}
			configs = append(configs, config)
		}

		// 进行查询
		for _, c := range configs {
			deviceConfig, err := configHTTPClient.GetDeviceConfig(
				context.Background(),
				&v1.GetDeviceConfigRequest{
					DeviceClassId: 0,
					DeviceId:      c.Id,
				},
			)
			if err != nil {
				t.Error(err)
				return
			}

			// 检查查询结果
			if !proto.Equal(&c, deviceConfig) {
				t.Errorf("Configuration query result error:%v %v", c, *deviceConfig)
			}
		}
	})

	// 2. 测试设备状态的批量查询功能
	t.Run("Test_BatchGetDeviceState", func(t *testing.T) {
		// 初始化测试用例
		var states []utilApi.TestedDeviceState
		for i := 0; i < 5; i++ {
			// 为了不影响之后的预警检测测试，这里写入不违反注册信息处填写的预警规则的设备状态
			state := utilApi.TestedDeviceState{
				Id: fmt.Sprintf("%s%d", t.Name(), i),
				// 注意这里写入的设备状态的时间应该为递增顺序，便于后续批量查询的检查
				Time:        timestamppb.New(time.Now().Add(time.Duration(i) * time.Second)),
				Voltage:     0,
				Current:     1000,
				Temperature: 0,
			}
			fields := map[string]float64{
				"voltage":     state.Voltage,
				"current":     state.Current,
				"temperature": state.Temperature,
			}
			err := saveState(redisClient, state.Id, &state, state.Time.AsTime().Unix(), fields)
			if err != nil {
				t.Fatal(err)
				return
			}
			states = append(states, state)
		}

		// 执行查询，并检查查询结果
		reply, err := warningDetectHTTPClient.BatchGetDeviceState(
			context.Background(),
			&v1.BatchGetDeviceStateRequest{
				DeviceClassId: 0,
				Start:         nil,
				End:           nil,
				Page:          0,
				Count:         0,
			},
		)
		if err != nil {
			t.Error(err)
			return
		}

		if len(reply.States) != len(states) {
			t.Error("State query result length error")
			return
		}

		for i, s := range states {
			// 检查查询结果
			if !proto.Equal(&s, reply.States[i]) {
				t.Errorf("State query result error:%v %v", s, *(reply.States[i]))
			}
		}
	})

	// 3. 测试关于设备状态的注册信息的查询功能
	t.Run("Test_GetRegisterInfo", func(t *testing.T) {
		for i, info := range registerInfo {
			stateRegisterInfo, err := warningDetectHTTPClient.GetDeviceStateRegisterInfo(
				context.Background(),
				&v1.GetDeviceStateRegisterInfoRequest{DeviceClassId: int64(i)},
			)
			if err != nil {
				t.Error(err)
				return
			}

			// 检查信息
			if !proto.Equal(&info, stateRegisterInfo) {
				t.Errorf("Register info query result error:%v %v", info, *stateRegisterInfo)
			}
		}
	})

	// 第四个测试和第五个测试中共用的设备状态信息
	var states []utilApi.TestedDeviceState

	// 4. 测试预警检测与多连接的预警推送功能
	t.Run("Test_WarningDetect", func(t *testing.T) {
		// 准备违反预警规则的设备状态信息
		for i := 0; i < 5; i++ {
			state := utilApi.TestedDeviceState{
				Id:          fmt.Sprintf("%s%d", t.Name(), i),
				Voltage:     1000,
				Current:     0,
				Temperature: 100,
			}
			states = append(states, state)
		}

		// 建立接收警告消息的channel组
		warningChannels := make([]chan *utilApi.Warning, 2)
		for i := 0; i < len(warningChannels); i++ {
			warningChannels[i] = make(chan *utilApi.Warning, len(states))
		}

		// 定义接收警告信息的协程函数
		worker := func(done <-chan struct{}, conn *websocket.Conn, warnings chan<- *utilApi.Warning) {
			defer conn.Close()
			for {
				select {
				case <-done:
					return
				default:
					warning := new(utilApi.Warning)
					err := conn.ReadJSON(warning)
					if err != nil {
						return
					}
					warnings <- warning
				}
			}
		}

		// 建立并保存ws连接，并启动协程接收信息
		done := make(chan struct{})
		var connections []*websocket.Conn
		// 注册清理函数
		t.Cleanup(func() {
			close(done)
			for _, c := range connections {
				c.Close()
			}
			for _, c := range warningChannels {
				close(c)
			}
		})

		for i := 0; i < len(warningChannels); i++ {
			conn, resp, err := websocket.DefaultDialer.Dial(
				"ws://localhost:8000/warnings/push", nil)
			if err != nil {
				t.Error(err)
				buffer := bytes.NewBuffer(make([]byte, 0, 1024))
				err := resp.Write(buffer)
				if err == nil {
					t.Errorf("%s", buffer.String())
				}
				return
			}
			connections = append(connections, conn)
			j := i
			go worker(done, conn, warningChannels[j])
		}

		// 将违反预警规则的设备状态信息写入
		now := time.Now()
		for i, s := range states {
			fields := map[string]float64{
				"voltage":     s.Voltage,
				"current":     s.Current,
				"temperature": s.Temperature,
			}
			s.Time = timestamppb.New(now.Add(time.Duration(i+5) * time.Second))
			err := saveState(redisClient, s.Id, &s, s.Time.AsTime().Unix(), fields)
			if err != nil {
				return
			}
		}

		// 等待一段时间，然后检查channel中接收到的警告消息
		time.Sleep(20 * time.Second)
		for _, c := range warningChannels {
			// 从channel接收每条警告的最大等待时间为10秒，
			// 若10秒内没读到期望的警告，则视作测试失败
			for _, s := range states {
				select {
				case <-time.After(10 * time.Second):
					t.Error("The receipt of warning information has timed out")
				case warning := <-c:
					if s.Id != warning.DeviceId {
						t.Errorf(
							"The received warning message does not match the expected:%v %v",
							s, *warning,
						)
					}
				}
			}
		}
	})

	// 5. 测试警告信息的批量查询功能
	t.Run("Test_BatchGetWarning", func(t *testing.T) {
		reply, err := warningDetectHTTPClient.BatchGetWarning(
			context.Background(),
			&v1.BatchGetWarningRequest{
				Start: nil,
				End:   nil,
				Page:  0,
				Count: 0,
			},
		)
		if err != nil {
			return
		}

		if len(reply.Warnings) != len(states) {
			t.Error("State query result length error")
			return
		}

		for i, s := range states {
			// 检查查询结果
			if reply.Warnings[i].DeviceId != s.Id {
				t.Errorf(
					"The received warning message does not match the expected:%v %v",
					s, *(reply.Warnings[i]),
				)
			}
		}

	})
}
