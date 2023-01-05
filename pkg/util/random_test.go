package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomCode(t *testing.T) {
	for i := 0; i < 30; i++ {
		time.Sleep(1)
		fmt.Println(RandomCode(6))
	}
}
