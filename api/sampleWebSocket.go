package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Client Connected")
	for {
		// Generate some random data
		data := []float64{rand.Float64(), rand.Float64(), rand.Float64()}

		// Send the data to the client
		err := conn.WriteJSON(data)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Data Written")
		// Wait for some time before sending the next data
		time.Sleep(1 * time.Second)
	}
}
