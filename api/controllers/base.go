package controllers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

type responseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}

var notifFailedGet string = "Failed Get Data"
var notifSuccessGet string = "Sukses Get Data"
var notifFailedUpdate string = "Failed Update Data"
var notifSuccessUpdate string = "Success Update Data"
var notifFailedDelete string = "Failed Delete Data"
var notifSuccessDelete string = "Success Delete Data"
var notifFailedInsert string = "Failed Insert Data"
var notifSuccessInsert string = "Success Insert Data"

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8180")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
