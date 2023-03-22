package rate_limiter

import (
	"context"
	"time"
)

const defaultWindowCount = 10

type SlidingWindow struct {
	windowSize   int64
	qps          int32           //窗口内的最大qps
	windowCount  int64           //子窗口个数
	windowsArray []int           //子窗口数组
	ctx          context.Context //控制限流协程的上下文
}

func (w *SlidingWindow) Start() {
	for {
		time.Sleep(time.Duration(w.windowSize / w.windowCount))
		w.windowsArray = append(w.windowsArray, 0)
		if len(w.windowsArray) > int(w.windowCount) {
			w.windowsArray = w.windowsArray[1:]

		}
	}
}

/*
func NewSlidingWindow(ctx context.Context, windowSize int64, qps int32) *SlidingWindow {
	windowsArray := make([]int, defaultWindowCount)
	window := &SlidingWindow{
		ctx:          ctx,
		windowSize:   windowSize,
		qps:          qps,
		windowCount:  defaultWindowCount,
		windowsArray: windowsArray,
	}

}
*/
