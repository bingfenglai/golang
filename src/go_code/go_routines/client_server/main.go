package main

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
			go run(do, req)

		case <-quit:
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
	req, quit := startServer(func(name string) string {
		return "hello! " + name
	})

	const N = 100

	var reqs [N]request

	for i := 0; i < N; i++ {
		r := reqs[i]

	}

}
