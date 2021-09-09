package model

import "sync"

// 修仙者等级
type levelLock struct {
	Lock sync.Mutex
	level      string
	levelValue int
}

func NewLevelLock(level string, levelValue int) *levelLock {
	return &levelLock{
		level:      level,
		levelValue: levelValue,
	}
}

func (recv *levelLock) SetLevel(level string) {

	recv.level  = level

}

