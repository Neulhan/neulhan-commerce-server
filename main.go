package main

import (
	"log"
	"neulhan-commerce-server/src/rest"
)

func main() {
	log.Println("START SERVER")
	log.Fatal(rest.RunAPI(":8000"))
}
