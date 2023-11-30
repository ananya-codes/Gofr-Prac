package main

import "gofr.dev/pkg/gofr"

func eventHandler(ctx *gofr.Context) (interface{}, error) {
	return "Hello World! from gofr event", nil
}
func main() {
	app := gofr.New()

	app.GET("/event", eventHandler)

	app.Start()
}
