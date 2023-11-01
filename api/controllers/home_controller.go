package controllers

import (
	"net/http"
	"asia-southeast2-bei-eipo-trial-apigee.cloudfunctions.net/ep_sleep_latency/api/responses"
	"time"
	"math/rand"
	"strconv"
)

type responseLatency struct {
	Message string          `json:"message"`
	Error   string          `json:"error"`
	Latency string 			`json:"latency"`
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	var time_latency int = rand.Intn(180)
	time.Sleep(time.Duration(time_latency) * time.Second)
	
	dataResponse := responseLatency{notifSuccessGet, "No Error", strconv.Itoa(time_latency)}

	responses.JSON(w, http.StatusOK, dataResponse)
}