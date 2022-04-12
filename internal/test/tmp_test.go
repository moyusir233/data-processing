package test

import (
	"context"
	"fmt"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/golang/protobuf/proto"
	"github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
	"testing"
	"time"
)

// 测试读写锁
func TestSync_RWMutex(t *testing.T) {
	lock := new(sync.RWMutex)
	done := make(chan struct{})
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					lock.RLock()
					fmt.Println("read")
					lock.RUnlock()
				}
			}
		}()
	}
	time.Sleep(10 * time.Millisecond)
	begin := time.Now()
	lock.Lock()
	fmt.Println("write")
	lock.Unlock()
	close(done)
	wg.Wait()
	fmt.Println(time.Since(begin))
}
func TestWriteInfluxdb(t *testing.T) {
	client := influxdb2.NewClient("http://localhost:8086", "L5Uo-KWQ1GvJ3gVV00Sc6TRrxTSnIVxiOOWcL5t1QFv42kP2cQMoPWPJSrWOWxv-t-4IffAA5Jkj9LXt34hVRQ==")
	defer client.Close()

	//var states []v1.DeviceState1
	//tags := map[string]string{"deviceClassID": "1"}
	//fields := map[string]float64{
	//	"tmp": 123.0,
	//}
	//now := time.Now()
	//
	//for i := 0; i < 50; i++ {
	//	// 为了不影响之后的预警检测测试，这里写入不违反注册信息处填写的预警规则的设备状态
	//	state := v1.DeviceState1{
	//		Id: fmt.Sprintf("%s%d", t.Name(), i),
	//		// 注意这里写入的设备状态的时间应该为递增顺序，便于后续批量查询的检查
	//		Time:        timestamppb.New(now.Add(time.Duration(i) * time.Minute)),
	//		Voltage:     0,
	//		Current:     1000,
	//		Temperature: 0,
	//	}
	//	err := saveState(client,
	//		"test", "test", state.Time.AsTime(), state.Id,
	//		fields, tags,
	//	)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	states = append(states, state)
	//}

	queryAPI := client.QueryAPI("test")
	flux := `from(bucket: "test")
				  |> range(start: -2h)
				  |> filter(fn: (r) => r["deviceClassID"] == "1")
				  |> group()
				  |> count()`

	tableResult, err := queryAPI.Query(context.Background(), flux)
	if err != nil {
		t.Fatal(err)
	}
	defer tableResult.Close()

	for tableResult.Next() {
		value := tableResult.Record().Value()
		fmt.Printf("value: %T %v\n", value, value)
		fmt.Println(tableResult.Record().String())
	}

}

func TestClearInfluxdb(t *testing.T) {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://gd-k8s-master01:30086", "test")
	// always close client at the end
	defer client.Close()

	deleteAPI := client.DeleteAPI()
	start := time.Now().UTC().Add(-1 * time.Hour)
	end := time.Now().UTC()
	predicate := fmt.Sprintf(`deviceClassID="%d"`, 1)
	for _, bucket := range []string{"test", "test-warning_detect", "test-warnings"} {
		deleteAPI.DeleteWithName(context.Background(), "test", bucket, start, end, predicate)
	}
}

func TestTmp(t *testing.T) {
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
			Start:           timestamppb.New(now.UTC()),
			End:             timestamppb.New(now.Add(5 * time.Minute).UTC()),
		},
	}

	warnings = append(
		warnings, copyAndAssign(warnings[0], "start", timestamppb.New(now.UTC())).(*utilApi.Warning))

	fmt.Printf("%v", warnings[1])
}
