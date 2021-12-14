package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(2)
	ch := make(chan string)
	ch1 := make(chan string)

	go sendData(ch)
	go sayHello(ch, ch1)

	<-ch1

}

func sendData(ch chan string) {

	ch <- "韩立"
	ch <- "厉飞羽"
	ch <- "张铁"
	ch <- "墨大夫"
	ch <- "南宫婉"
	ch <- "六道传人"
	ch <- "董萱儿"
	ch <- "EOF"
}

func sayHello(ch, ch1 chan string) {

	for {
		input := <-ch

		if input != "EOF" {
			println("hello !", input)
		} else {
			break
		}

	}

	ch1 <- "EOF"
}
