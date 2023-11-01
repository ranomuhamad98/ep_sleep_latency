package api

import (
	"fmt"
	"asia-southeast2-bei-eipo-trial-apigee.cloudfunctions.net/ep_sleep_latency/api/controllers"
)

var server = controllers.Server{}

func Run() {	
	fmt.Println("We are getting the env values")
	server.Initialize()
	server.Run(":8180")
}
