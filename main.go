package main

import (
	"log"

	"github.com/DanilKlochkov/golang-graphql/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
