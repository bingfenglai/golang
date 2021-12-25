package interval

import (
	"errors"
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

// 获取序列的下一个区间
func (receiver *Interval) GetNextInterval() (min, max int, err error) {

	receiver.lock.Lock()
	defer receiver.lock.Unlock()

	if receiver.currentVale >= receiver.maxValue {
		return 0, 0, errors.New("序列号已耗尽")
	}

	min = receiver.currentVale
	max = receiver.currentVale + receiver.step
	if max > receiver.maxValue {
		max = receiver.maxValue
	}
	receiver.currentVale = max + 1
	return min, max, nil
}
