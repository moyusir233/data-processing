package test

import (
	"fmt"
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
