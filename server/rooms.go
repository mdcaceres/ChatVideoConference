package server

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init initialises the RoomMap struct
func (r *RoomMap) Init() {
	//r.Mutex.Lock()
	//defer r.Mutex.Unlock()
	r.Map = make(map[string][]Participant)
}

// Get will return the array of participants in the room
func (r *RoomMap) Get(roomId string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomId]
}

// CreateRoom generate a unique room Id a return it
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	var lettersPool = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	idByte := make([]rune, 8)

	for i := range idByte {
		idByte[i] = lettersPool[random.Intn(len(lettersPool))]
	}

	roomId := string(idByte)

	r.Map[roomId] = []Participant{}

	return roomId
}

func (r *RoomMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Printf("Inserting into room id %s", roomId)
	r.Map[roomId] = append(r.Map[roomId], p)
}

func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}
