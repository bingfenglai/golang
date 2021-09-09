package model


// 修仙者等级
type Level struct {
	level      string
	levelValue int
}

func NewLevel(level string, levelValue int) Level {
	return Level{
		level:      level,
		levelValue: levelValue,
	}
}


// 获取等级描述
func (recv Level) Level() string{
	return recv.level
}

func (recv *Level) SetLevel(level string) {
	recv.level = level
}

func (recv *Level) LevelName() string{
	return recv.level
}

