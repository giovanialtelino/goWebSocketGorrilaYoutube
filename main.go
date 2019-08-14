package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giovanialtelino/youtube-stats/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(w, "%+v\n", err)
	}

	go websocket.Writer(ws)

}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtube Subscriber Monitor")
	setupRoutes()
}
