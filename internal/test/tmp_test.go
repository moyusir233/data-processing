package test

import (
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/influxdata/influxdb-client-go/v2"
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
	client := influxdb2.NewClient("http://localhost:8086", "BBQcGQrP1crCOe_alluyoagCjOaRo233oCvlBvYdcEoc7DM8MgzgF7YPzLkDxZ_LA92UsAJ0LiuvV2KwyQ2qfw==")
	defer client.Close()

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
		err := saveState(client,
			"test", "test", state.Time.AsTime(), state.Id,
			fields, tags,
		)
		if err != nil {
			t.Fatal(err)
		}
		states = append(states, state)
	}
	//
	//queryAPI := client.QueryAPI("test")
	//flux := `from(bucket: "test")
	//			  |> range(start: %d, stop: %d)
	//			  |> filter(fn: (r) => r["deviceClassID"] == "1")
	//			  |> group(columns: ["deviceClassID", "_measurement", "_time"])`
	//
	//tableResult, err := queryAPI.Query(context.Background(),
	//	fmt.Sprintf(flux, now.UTC().Unix(), now.Add(5*time.Second).UTC().Unix()))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer tableResult.Close()
	//
	//for tableResult.Next() {
	//	table := tableResult.Record().ValueByKey("table")
	//	fmt.Printf("table: %T %v\n", table, table)
	//	fmt.Println(tableResult.Record().String())
	//}

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
	timestamp := timestamppb.New(time.Now())
	marshal, err := encoding.GetCodec("json").Marshal(timestamp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshal))
}