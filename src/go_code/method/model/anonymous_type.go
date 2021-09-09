package model



// 修仙者
type immortal2 struct {
	name   string
	age    int
	gender string
	Level
	lingGen
}

func NewImmortal2(age int, name, gender string,levelName string,levelValue int,lingGenNames...string) *immortal2 {
	return &immortal2{
		name:   name,
		age:    age,
		gender: gender,
		Level:  Level{levelName,levelValue},
		lingGen: lingGen{linGenNames: lingGenNames},
	}
}
