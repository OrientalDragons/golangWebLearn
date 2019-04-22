package main

import (
	"log"
	"net/http"

	"go_code/goMyWeb0.0.3/routes"
)

func main() {

	err := http.ListenAndServe(":80", routes.GetRoute())
	log.Print(err)
}
