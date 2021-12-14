package main

func main() {
	ch := make(chan string)
	go sendData(ch)
	go receivingData(ch)

}

func sendData(ch chan<- string) {

}

func receivingData(ch <-chan string) {

}
