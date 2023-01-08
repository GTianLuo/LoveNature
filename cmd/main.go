package main

import (
	"lovenature/log"
	"lovenature/routes"
)

func main() {

	r := routes.NewRouter()
	if err := r.Run(":9999"); err != nil {
		log.Error(err)
	}

}
