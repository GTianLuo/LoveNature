package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	s := []string{"abdd", "shsjhs", "hdhiuhjdiu", "djuhduir"}
	sjson, _ := json.Marshal(s)
	fmt.Println(string(sjson))
	var s1 []string
	if err := json.Unmarshal(sjson, &s1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)
}
