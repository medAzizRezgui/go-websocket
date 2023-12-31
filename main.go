package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		//	read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home HTTP")
}
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	// Listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)

}
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
