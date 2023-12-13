package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	// register route greet
	app.GET("/health", func(ctx *gofr.Context) (interface{}, error) {
		//stats := ctx.DB().Stats()
		//fmt.Println("stats ", stats)
		return "Health OK!!", nil
	})
	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
