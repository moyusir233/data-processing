package test

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
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
	client := influxdb2.NewClient("http://localhost:8086", "Sub63Yc3aYwi0dqc8HBJl6P_Ev4X-ygyMumq_CEEWgNBvQ0OcmqNHAQASFYxOI5Ai02vtzTBVT80_yWp0QhwRA==")
	defer client.Close()

	writeAPI := client.WriteAPI("test", "test")
	defer writeAPI.Flush()
	deviceClassID := "0"
	now := time.Now().UTC().Add(-5 * time.Minute)
	for i := 0; i < 300; i++ {
		for j := 0; j < 2; j++ {
			deviceID := fmt.Sprintf("device_%d", j)
			point := write.NewPointWithMeasurement(deviceID).
				AddTag("deviceClassID", deviceClassID).
				AddField("current", i).SetTime(now)
			writeAPI.WritePoint(point)
		}
		now = now.Add(time.Second)
	}
}
