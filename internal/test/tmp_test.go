package test

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
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
	client := influxdb2.NewClient("http://gd-k8s-master01:30086", "test")
	defer client.Close()

	deleteAPI := client.DeleteAPI()
	start := time.Now().UTC().Add(-1 * time.Hour)
	end := time.Now().UTC()
	predicate := fmt.Sprintf(`deviceClassID="%d"`, 1)
	for _, bucket := range []string{"test", "test-warning_detect", "test-warnings"} {
		deleteAPI.DeleteWithName(context.Background(), "test", bucket, start, end, predicate)
	}
}
