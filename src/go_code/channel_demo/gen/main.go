package main

import (
	fmt "fmt"
	"sync"
)

type Interval struct {
	maxValue    int
	currentVale int
	step        int
	lock        sync.Mutex
}

func NewInterval(maxValue, step int) *Interval {

	return &Interval{
		maxValue:    maxValue,
		currentVale: 1,
		step:        step,
	}

}

func (receiver *Interval) getNextInterval() (min, max int) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	if receiver.currentVale >= receiver.maxValue {
		panic("序列号已耗尽")
	}

	min = receiver.currentVale
	max = receiver.currentVale + receiver.step
	if max > receiver.maxValue {
		max = receiver.maxValue
	}
	receiver.currentVale = max + 1
	return min, max
}

var resume chan int

var interval = NewInterval(99999, 1000)

func intergers() chan int {
	yeild := make(chan int)
	min, max := interval.getNextInterval()

	go func() {
		for {

			if min < max {
				yeild <- min
				min++
			} else {
				min, max = interval.getNextInterval()
			}

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
	for {
		fmt.Println(fmt.Sprintf("%05d", generateInteger()))
	}

	//for  {
	//	fmt.Println(interval.getNextInterval())
	//
	//}

}
