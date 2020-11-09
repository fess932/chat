package main

import (
	"flag"
	"log"

	"github.com/kataras/iris/v12/middleware/logger"

	"github.com/gorilla/websocket"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

var addr = flag.String("addr", ":8000", "http server address")

func main() {
	flag.Parse()

	log.Println("prepare...")

	app := iris.New()
	app.UseRouter(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}))
	app.UseRouter(logger.New())

	// setup rest path
	rest(app)

	log.Println("start server:", *addr)
	log.Fatal(app.Listen(*addr))

}

func rest(app *iris.Application) {
	app.Get("/pinn", ping)

	app.Get("/v1/ws", wsHandler)
}

func ping(_ iris.Context) {
	return
}

func wsHandler(ctx iris.Context) {
	//conn, _ := websocket.Upgrade(r, w)
	//ch := NewChannel(conn)
	////...
	ws := websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}

	conn, err := ws.Upgrade(ctx.ResponseWriter(), ctx.Request(), ctx.Request().Header)
	if err != nil {
		//log.Println(err)
		ctx.StopWithError(400, err)
		return
	}

	conn.PingHandler()

	log.Println("handle func... ")
}
