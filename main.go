package main

import (
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()
	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
