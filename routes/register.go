package routes

import (
	"net/http"

	"go_code/goMyWeb0.0.3/controller"
)

var regularRoute = map[string]func(http.ResponseWriter, *http.Request){
	// "/static/.*": controller.TestNext,
	"/static/":   controller.Funtest,
	"/uploap/.*": controller.Funtest2,
}

var absoluteRoute = map[string]func(http.ResponseWriter, *http.Request){

	"/test/p": controller.Funtest3,
}
