package main

import (
	"chat-server/config"
	"chat-server/internal/server/wsServer"
	"flag"
	"log"

	"github.com/kataras/iris/v12/middleware/logger"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

var addr = flag.String("addr", ":8000", "http server address")

func main() {
	flag.Parse()

	config.CreateRedisClient()

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
	app.Get("/ping", ping)

	// ws conf
	wsHandler(app)
}

func ping(_ iris.Context) {
	return
}

func wsHandler(app *iris.Application) {

	ws := wsServer.NewWebsocketServer()
	go ws.Run()

	app.Get("/v1/ws", func(ctx iris.Context) {
		if err := ws.ServeWS(ctx.ResponseWriter(), ctx.Request()); err != nil {
			log.Println(err)
			ctx.StopWithError(400, err)
		}
	})
}
