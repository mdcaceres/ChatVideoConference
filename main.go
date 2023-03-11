package main

import (
	"github.com/mdcaceres/chat/server"
	"log"
	"net/http"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting server on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
