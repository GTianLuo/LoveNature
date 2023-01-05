package util

import (
	"fmt"
	"testing"
)

func TestSnowflake_NextVal(t *testing.T) {
	m := make(map[int64]bool, 100000)
	/*
		for i := 0; i < 100; i++ {
			fmt.Println(s.NextVal())
		}*/

	for i := 0; i < 100000; i++ {
		v := IdGenerator.NextVal()
		fmt.Println(v)
		//fmt.Println(v)
		if m[v] == true {
			//fmt.Println(IdGenerator.timeStamp)
			return
		}
		m[v] = true
	}
	fmt.Println("成功生成100000个id，无重复")

}

func TestGetNickName(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetNickName())
	}
}
