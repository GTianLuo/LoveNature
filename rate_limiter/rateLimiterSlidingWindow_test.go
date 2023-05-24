package rate_limiter

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	window := NewSlidingWindow(1000, 2, context.Background())
	var wg sync.WaitGroup
	var count int32
	start := time.Now().Second()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				//time.Sleep(300 * time.Millisecond)
				if window.TryQuery() {
					atomic.AddInt32(&count, 1)
					fmt.Println("Success")
				} else {
					fmt.Println("Fail")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("花费时间:", time.Now().Second()-start)
	fmt.Println("执行次数:", count)
}
