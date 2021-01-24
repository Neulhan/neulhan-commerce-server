package main

import (
	"log"
	"neulhan-commerce-server/src/config"
	"neulhan-commerce-server/src/rest"
	"os"
)

func main() {
	log.Printf("[%s] START SERVER ON %s", os.Getenv("GIN_MODE"), config.GetEnv("PORT"))
	log.Fatal(rest.RunAPI(config.GetEnv("PORT")))
}
