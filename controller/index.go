package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"go_code/goMyWeb0.0.3/model"
)

func init() {}

//Funtest 测试
func Funtest(w http.ResponseWriter, r *http.Request) {
	log.Println("-----------------------------------------")
	log.Println("something request:" + r.URL.Path)
	log.Println("request refernce:" + r.Referer())
	setStaticPath("/static/", w, r)
	return

}
func Funtest2(w http.ResponseWriter, r *http.Request) {
	conn := model.ConnDb()
	stringt := model.TestDb(conn)
	fmt.Println(stringt)
	io.WriteString(w, stringt) //如果write byte类型数据前端会显示下载文件
	return
}
func Funtest3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test3")
	io.WriteString(w, "test3")
	return
}
