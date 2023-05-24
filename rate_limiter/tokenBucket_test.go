package rate_limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	bucket := NewTokenBucket(2, 2, 2)

	for i := 0; i <= 30; i++ {
		time.Sleep(100 * time.Millisecond)
		if bucket.TryAcquired() {
			fmt.Println("Success")
		} else {
			fmt.Println("Fail")
		}
	}

}
