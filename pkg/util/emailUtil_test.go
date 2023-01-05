package util

import (
	"fmt"
	"testing"
)

func TestSendCode(t *testing.T) {
	SendCode("123456")
}

func TestVerifyEmailFormat(t *testing.T) {
	/*
		fmt.Println(VerifyEmailFormat("2985496686"))
		fmt.Println(VerifyEmailFormat("2985496686@qq.com"))
		fmt.Println(VerifyEmailFormat("2985496686@qq"))
	*/
	fmt.Println(VerifyPasswordFormat("pp;';;'87855455"))
	fmt.Println(VerifyPasswordFormat("87855455"))
	fmt.Println(VerifyPasswordFormat("8RE"))
	fmt.Println(VerifyPasswordFormat("8REsdsdsdsds"))
	fmt.Println(VerifyPasswordFormat("8REsdsdsdsdserdf"))
	fmt.Println(VerifyPasswordFormat("8REsdsdsdsdserdfd"))
	fmt.Println(VerifyPasswordFormat("8855555dD"))

}
