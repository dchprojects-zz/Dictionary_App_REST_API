package main

import (
	"goproj/app"
	"goproj/config"
)

func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")

}
