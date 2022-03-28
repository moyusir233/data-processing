package test

import (
	"bytes"
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/data"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
	"strings"
	"sync"
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
			Fields: nil,
		},
		{
			// 以v1.DeviceState1为模板进行定义
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
					Type: utilApi.Type_INT64,
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
	redisClient, cleanup, err := data.NewRedisData(bootstrap.Data)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		// 测试结束时清空数据库
		redisClient.FlushDB(context.Background())
		cleanup()
	})

	// 初始化用于写入测试用例信息的influx客户端
	influxClient, cleanup2, err := data.NewInfluxdbData(bootstrap.Data)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		for _, bucket := range []string{"test", "test-warning_detect", "test-warnings"} {
			clearInfluxdb(influxClient, "test", bucket, 1)
		}
		cleanup2()
	})

	// 启动服务器，并实例化测试所需的http客户端
	configHTTPClient, warningDetectHTTPClient := StartDataProcessingServer(t, bootstrap, registerInfo)

	// 1. 测试设备的配置查询功能
	t.Run("Test_GetDeviceConfig", func(t *testing.T) {
		// 初始化测试数据，并写入到数据库中
		var configs []v1.DeviceConfig1
		for i := 0; i < 5; i++ {
			config := v1.DeviceConfig1{
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
			deviceConfig, err := configHTTPClient.GetDeviceConfig1(
				context.Background(),
				&v1.GetDeviceConfigRequest{
					DeviceId: c.Id,
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
		var states []v1.DeviceState1
		tags := map[string]string{"deviceClassID": "1"}
		now := time.Now()
		for i := 0; i < 5; i++ {
			// 为了不影响之后的预警检测测试，这里写入不违反注册信息处填写的预警规则的设备状态
			state := v1.DeviceState1{
				Id: fmt.Sprintf("%s%d", t.Name(), i),
				// 注意这里写入的设备状态的时间应该为递增顺序，便于后续批量查询的检查
				Time:        timestamppb.New(now.Add(time.Duration(i) * time.Second)),
				Voltage:     0,
				Current:     1000,
				Temperature: 0,
			}
			fields := map[string]float64{
				"tmp": 123.0,
			}
			err := saveState(influxClient,
				"test", "test", state.Time.AsTime(), state.Id,
				fields, tags,
			)
			if err != nil {
				t.Fatal(err)
			}
			states = append(states, state)
		}

		// 执行查询，并检查查询结果
		reply, err := warningDetectHTTPClient.BatchGetDeviceStateInfo(context.Background(),
			&v1.BatchGetDeviceStateRequest{
				DeviceClassId: 1,
				Start:         timestamppb.New(now),
				End:           timestamppb.New(now.Add(5 * time.Second)),
			},
		)
		if err != nil {
			t.Error(err)
			return
		}

		if len(reply.States) != len(states) {
			t.Errorf("State query result length error:%d %d", len(reply.States), len(states))
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
	var states []v1.DeviceState1

	// 4. 测试预警检测与多连接的预警推送功能
	t.Run("Test_WarningDetect", func(t *testing.T) {
		// 准备违反预警规则的设备状态信息
		for i := 0; i < 5; i++ {
			state := v1.DeviceState1{
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
			warningChannels[i] = make(chan *utilApi.Warning, len(states)*3)
		}

		// 定义接收警告信息的协程函数
		worker := func(
			done <-chan struct{},
			wg *sync.WaitGroup,
			conn *websocket.Conn,
			warnings chan<- *utilApi.Warning) {
			wg.Add(1)
			defer func() {
				conn.Close()
				wg.Done()
			}()

			var retry int

			for {
				select {
				case <-done:
					return
				default:
					warning := new(utilApi.Warning)
					err := conn.ReadJSON(warning)
					// 由于从连接中读取信息会造成协程阻塞，
					// 且可能是因为done关闭，然后关闭了ws连接而产生的错误，
					// 因此当首次发生错误时，先continue，去判断done是否关闭
					if err != nil {
						retry++
						if retry > 1 {
							t.Error(err)
							return
						} else {
							continue
						}
					}
					warnings <- warning
				}
			}
		}

		// 建立并保存ws连接，并启动协程接收信息
		done := make(chan struct{})
		wg := new(sync.WaitGroup)
		var connections []*websocket.Conn
		// 注册清理函数
		t.Cleanup(func() {
			// 首先通知协程进行关闭，然后关闭前端连接，然后等待协程结束，最后关闭channel以避免协程写入关闭的channel
			close(done)
			for _, c := range connections {
				c.Close()
			}
			wg.Wait()
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
			go worker(done, wg, conn, warningChannels[j])
		}

		// 将违反预警规则的设备状态信息写入
		now := time.Now()
		tags := map[string]string{"deviceClassID": "1"}
		for i, s := range states {
			fields := map[string]float64{
				"voltage":     s.Voltage,
				"current":     s.Current,
				"temperature": s.Temperature,
			}
			s.Time = timestamppb.New(now.Add(time.Duration(i+5) * time.Second))
			err := saveState(influxClient, "test", "test",
				s.Time.AsTime(), s.Id, fields, tags,
			)
			if err != nil {
				return
			}
		}

		// 等待一段时间，然后通过计数检查channel中接收到的警告消息
		time.Sleep(15 * time.Second)
		count := make(map[string]int, len(states))
		for _, c := range warningChannels {
			// 从channel接收每条警告的最大等待时间为3秒，
			// 若3秒内没读到期望的警告，则视作测试失败
			for i := 0; i < len(states); i++ {
				// 每条设备状态信息对应着三个预警字段，即三条警告信息
				for j := 0; j < 3; j++ {
					select {
					case <-time.After(3 * time.Second):
						t.Error("The receipt of warning information has timed out")
					case warning := <-c:
						count[warning.DeviceId]++
					}
				}
			}
		}

		// 检查是否每项设备状态信息都产生了3条警告记录
		if len(count) != len(states) {
			t.Errorf(
				"The received warning message count does not match the expected:%d %d",
				len(count), len(states),
			)
		}
		for _, v := range count {
			// 由于每个state有三项预警字段，因此每个state对应着3条警告信息，
			// 而由于推送时会将相同的警告消息扇出到所有的前端连接中，
			// 因此每个state对应的警告信息计数为前端连接数*3
			if v != 3*len(connections) {
				t.Errorf(
					"The received warning message count does not match the expected:%d %d",
					v, 3*len(connections),
				)
			}
		}
	})

	// 5. 测试警告信息的批量查询功能
	t.Run("Test_BatchGetWarning", func(t *testing.T) {
		reply, err := warningDetectHTTPClient.BatchGetWarning(
			context.Background(),
			&v1.BatchGetWarningRequest{
				Past: durationpb.New(5 * time.Minute),
			},
		)
		if err != nil {
			return
		}

		if len(reply.Warnings) != 3*len(states) {
			t.Errorf("State query result length error:%d %d", len(reply.Warnings), 3*len(states))
			return
		}

		// 对查询得到的警告信息按照id的降序进行排序，然后与states中保存的设备状态进行比较,
		// 以确定每项设备状态信息产生了三条警告信息
		sort.Slice(reply.Warnings, func(i, j int) bool {
			return reply.Warnings[i].DeviceId < reply.Warnings[j].DeviceId
		})

		for i, s := range states {
			// 检查查询结果
			// 由于警告信息中返回的是完整的设备key，即<用户id>:<设备类别号>:<设备id>，
			// 因此这里利用后缀来判断id警告信息是否正确
			for j := i * 3; j < (i+1)*3; j++ {
				if !strings.HasSuffix(reply.Warnings[j].DeviceId, s.Id) {
					t.Errorf(
						"The received warning message does not match the expected:%s %s",
						s.Id, reply.Warnings[j].DeviceId,
					)
				}
			}
		}

	})
}
