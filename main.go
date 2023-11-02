package main

import (
	"fmt"
	"encoding/json"
	"net/http"	
	"time"
	"math/rand"
	"strconv"
	"log"
	"gorilla/mux"
)

type responseLatency struct {
	Message string          `json:"message"`
	Error   string          `json:"error"`
	Latency string 			`json:"latency"`
}
type responseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}
type Server struct {
	Router *mux.Router
}

var notifFailedGet string = "Failed Get Data"
var notifSuccessGet string = "Success Get Data"
var notifFailedUpdate string = "Failed Update Data"
var notifSuccessUpdate string = "Success Update Data"
var notifFailedDelete string = "Failed Delete Data"
var notifSuccessDelete string = "Success Delete Data"
var notifFailedInsert string = "Failed Insert Data"
var notifSuccessInsert string = "Success Insert Data"

var server = Server{}

func main() {
	// api.Run()
	fmt.Println("We are getting the env values")
	server.Initialize()
	server.Run(":8180")
}

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", SetMiddlewareJSON(s.Home)).Methods("GET")
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8180")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	var time_latency int = rand.Intn(180)
	time.Sleep(time.Duration(time_latency) * time.Second)
	
	dataResponse := responseLatency{notifSuccessGet, "No Error", strconv.Itoa(time_latency)}

	JSON(w, http.StatusOK, dataResponse)
}