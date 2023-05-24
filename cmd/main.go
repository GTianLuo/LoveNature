package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		r := routes.NewRouter()
		if err := r.Run(":9999"); err != nil {
			log.Error(err)
		}*/

	for i := 0; i < 2000000; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				fmt.Println(".")
			}
		}()
	}
	time.Sleep(1000000 * time.Second)
}
