package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type a struct {
	X int
	Y *string
}

type B struct {
	Y a
	Z []int
}

func TestStructToMap(t *testing.T) {
	s := "3444"
	A := &a{
		X: 1,
		Y: &s,
	}
	b := &B{
		Y: *A,
		Z: []int{1, 2, 3, 4, 5, 6, 76},
	}
	m := StructToMap(b)
	fmt.Println(m)
	bb := &B{}
	MapToStruct(m, bb)
	fmt.Println(bb)
}

func TestJson(t *testing.T) {
	s := "3444"
	A := &a{
		X: 1,
		Y: &s,
	}
	b := &B{
		Y: *A,
		Z: []int{1, 2, 3, 4, 5, 6, 76},
	}
	marshal, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))
	bb := &B{}
	fmt.Println(json.Unmarshal(marshal, bb))
}

func TestReflect(t *testing.T) {
	s := "12122"
	a := a{
		X: 1,
		Y: &s,
	}
	fmt.Println(reflect.ValueOf(a).FieldByName("Y").Type().String())
	field := reflect.ValueOf(a).FieldByName("Y").Elem()

	field.Set(reflect.ValueOf("sssssss"))
	fmt.Println(a)
}
