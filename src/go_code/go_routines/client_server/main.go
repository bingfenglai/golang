package main

import (
	"strconv"
	"time"
)

// 请求
type request struct {
	args   string
	replyc chan string
}

type sayHello func(name string) string

func run(do sayHello, req *request) {
	req.replyc <- do(req.args)
}

func server(do sayHello, service chan *request, quit chan int) {
	for {
		select {
		case req := <-service:
			println("收到请求", req.args)
			go run(do, req)

		case <-quit:
			println("收到退出指令")
			return
		}

	}
}

func startServer(do sayHello) (seivice chan *request, quit chan int) {
	seivice = make(chan *request)
	quit = make(chan int)
	go server(do, seivice, quit)
	return seivice, quit

}

func main() {
	service, quit := startServer(func(name string) string {
		return "hello! " + name
	})

	const N = 100

	var reqs [N]request

	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.args = strconv.Itoa(i) + " name"
		req.replyc = make(chan string)
		service <- req

	}

	for i := 0; i < N; i++ {
		req := &reqs[i]
		s := <-req.replyc
		println(s)
	}
	quit <- 1
	time.Sleep(1 * 1e9)

}
