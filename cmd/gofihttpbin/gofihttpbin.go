package main

import (
	"log"

	"github.com/werdes72/gofihttpbin/pkg/gofihttpbin"
)

func main() {
	app := gofihttpbin.NewApp()

	log.Fatal(app.Listen(":8080"))
}
