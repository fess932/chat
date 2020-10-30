package main

import (
	"chat-server/config"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8000", "http server address")

func main() {
	flag.Parse()

	config.InitDB()

	wsServer := NewWebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Println("log....")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
