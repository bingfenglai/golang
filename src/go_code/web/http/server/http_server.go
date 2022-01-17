package main

import (
	"go_code/web/http/constants"
	"net/http"
	"strings"
)

var handlerMap = map[string]http.HandlerFunc{}

func main() {

	handlerMap["/sayHello"] = SayHelloHandler
	handlerMap["/goodbye"] = SayGoodbyeHandler

	http.HandleFunc(constants.CONTEXT_PATH, BaseHandler)

	err := http.ListenAndServe(constants.ADDR, nil)

	if err != nil {
		println("启动http服务器错误\n", err.Error())
	}

}

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	uri := strings.Split(r.RequestURI, "?")[0]

	println("请求URI ", uri)
	//println(i)

	handlerFunc, ok := handlerMap[uri]
	if ok {
		handlerFunc(w, r)
	} else {

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("请求资源不存在"))

	}
}

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {

	// 获取get请求参数
	name := r.URL.Query().Get("name")

	if name != "" {
		_, _ = w.Write([]byte("hello ! " + name))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("缺少参数name"))
	}
}

func SayGoodbyeHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name != "" {
		_, _ = w.Write([]byte("goodbye ! " + name))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("缺少参数name"))
	}
}
