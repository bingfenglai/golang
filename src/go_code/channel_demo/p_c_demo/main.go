package main

func main() {

	Consume(Produce(10))
}

func Produce(size int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < size; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Consume(ch <-chan int) {
	for i := range ch {
		println("收到数据", i)
	}
}
