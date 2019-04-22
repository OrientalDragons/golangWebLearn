package controller

import (
	"log"
	"net/http"
	"os"
	"path"
)

//StaticFn xxx
func setStaticPath(dir string, w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	dirPath := path.Join(wd, dir)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(dirPath)))
	staticHandler.ServeHTTP(w, r)
}
