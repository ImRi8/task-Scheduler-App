package main

import (
	"Task-scheduler-App/Handlers"
	"gofr.dev/pkg/gofr"
)

const (
	TASK_BACKEND = "/task"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	// register route greet
	app.GET(TASK_BACKEND+"/health", func(ctx *gofr.Context) (interface{}, error) {
		//stats := ctx.DB().Stats()
		//fmt.Println("stats ", stats)
		return "Health OK!!", nil
	})

	app.GET(TASK_BACKEND+"/getTaskById", Handlers.GetTaskByIdHandler)

	//app.GET("/")
	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
