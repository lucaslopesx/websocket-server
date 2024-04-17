package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var PORT *uint;
var upgrader = websocket.Upgrader{}
var clients []websocket.Conn

func init() {
	PORT = flag.Uint("port", 8080, "set port number")
	
	flag.Parse()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	clients = append(clients, *conn)

	defer conn.Close()
	
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message:", err)
			return
		}

		log.Println("message: ", string(message))

		for _, client := range clients {
			err = client.WriteMessage(mt, message)
			if err != nil {
				log.Println("write message:", err)
				return
			}
		}
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Handler() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/client", Client)
	http.HandleFunc("/echo", echo)
}

func main() {
	fmt.Printf("Starting go server at port %d.", *PORT)

	Handler()

	addr := fmt.Sprintf(":%d", *PORT)

	http.ListenAndServe(addr, nil)
}
