package main

import (
	"log"

	"github.com/werdes72/gofihttpbin/pkg/gofihttpbin"
)

func main() {
	app := gofihttpbin.NewApp("./web/static/")

	log.Fatal(app.Listen(":8080"))
}
