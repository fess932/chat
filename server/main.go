package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/rs/cors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer)

	go client.writePump()
	go client.readPump()

	wsServer.register <- client

	fmt.Println("New Client joined the hub!")
	fmt.Println(client)
}

func main() {

	app := iris.Default()
	//// Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.WrapRouter(c.ServeHTTP)

	wsServer := NewWebsocketServer()
	go wsServer.Run()

	app.Any("/ws", func(ctx *context.Context) {
		ServeWs(wsServer, ctx.ResponseWriter(), ctx.Request())
	})

	app.Listen(":8080")
}
