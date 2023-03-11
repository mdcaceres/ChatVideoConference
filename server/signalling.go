package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var AllRooms RoomMap

// CreateRoomRequestHandler CreateRoomHandler Create a Room and return roomId
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomId := AllRooms.CreateRoom()

	type resp struct {
		RoomId string `json:"room_id"`
	}

	log.Println(AllRooms.Map)

	json.NewEncoder(w).Encode(resp{RoomId: roomId})
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast
		for _, client := range AllRooms.Map[msg.RoomId] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}

// JoinRoomRequestHandler CreateRoomRequestHandler JoinRoomRequestHandler will join the client in a particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("roomId")

	if roomId == "" {
		log.Println("roomId missing in URL parameters")
		return
	}

	ws, err := upgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("web Socket upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomId, false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("read error:", err)
		}
		msg.Client = ws
		msg.RoomId = roomId

		broadcast <- msg
	}
}
