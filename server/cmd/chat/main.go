package main

import (
	"chat-server/internal/wsClient"
	"flag"
	"log"

	"github.com/kataras/iris/v12/middleware/logger"

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
	app.Get("/ping", ping)

	// ws conf
	wsHandler(app)
}

func ping(_ iris.Context) {
	return
}

func wsHandler(app *iris.Application) {
	hub := wsClient.NewHub()
	go hub.Run()



	app.Get("/v1/ws", func(ctx iris.Context) {
		if err := wsClient.ServeWS(hub, ctx.ResponseWriter(), ctx.Request()); err != nil {
			log.Println(err)
			ctx.StopWithError(400, err)
		}
	})
}

