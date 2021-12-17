package main

import (
	fmt "fmt"
	"sync"
)

type Interval struct {
	maxValue    int
	currentVale int
	step        int
	sync.Locker
}

func (receiver *Interval) getNextInterval() (min, max int) {
	receiver.Lock()
	defer receiver.Unlock()
	min = receiver.currentVale
	max = receiver.currentVale + receiver.step
	return min, max
}

var resume chan int

func intergers() chan int {
	yeild := make(chan int)
	count := 0

	go func() {
		for {
			yeild <- count
			count++

		}
	}()

	return yeild
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = intergers()

	s := fmt.Sprintf("%05d", 3)
	fmt.Println(s)
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())

}
