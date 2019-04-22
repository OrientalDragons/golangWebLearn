package routes

import (
	"net/http"
	"regexp"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", regularHandlerFunc)
	// mux.HandleFunc("/", controller.Funtest)
	for rt, fn := range absoluteRoute {
		mux.HandleFunc(rt, fn)
	}
}

func regularHandlerFunc(w http.ResponseWriter, r *http.Request) {
	found := false
	url := r.URL.Path
	for rt, fn := range regularRoute {
		if match, _ := regexp.MatchString(rt, url); match {

			if url != "/static/A.jpg" { //filter test
				found = true
				fn(w, r)
				return
			}

		}
	}
	if found == false {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

}

//GetRoute : 返回mux控制器
func GetRoute() *http.ServeMux {
	return mux
}
