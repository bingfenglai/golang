package main

import (
	"runtime"
	"strconv"
	"time"
)

// 请求
type request struct {
	args   string
	replyc chan string
}

type sayHello func(name string) string

// 模拟调度具体的业务方法，并通过通道返回结果
func run(do sayHello, req *request) {
	req.replyc <- do(req.args)
}

// 模拟服务器应用
func server(do sayHello, service chan *request, quit chan int) {
	defer func() {
		close(quit)
		for {
			if cap(service) == 0 {
				close(service)
				println("程序退出")
				return
			}

		}
	}()
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

// 启动服务器方法
// resp: 请求通道、退出信号通道
func startServer(do sayHello) (seivice chan *request, quit chan int) {
	seivice = make(chan *request)
	quit = make(chan int)
	go server(do, seivice, quit)
	return seivice, quit

}

func main() {
	runtime.GOMAXPROCS(4)
	service, quit := startServer(func(name string) string {
		return "hello! " + name
	})

	const N = 100

	var reqs [N]request

	// 初始化N个请求实例，并对服务器发器请求
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.args = strconv.Itoa(i) + " name"
		req.replyc = make(chan string)

		// 模拟接收响应结果
		go func() {
			s := <-req.replyc
			println("接收响应：", s)

		}()
		service <- req
		if i == N-1 {
			quit <- 1
		}

	}

	time.Sleep(2 * 1e9)

}
