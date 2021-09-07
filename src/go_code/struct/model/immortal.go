package model

// 修仙者
type Immortal struct {
	Name   string
	Age    int
	Gender string
}

func getImmortal(age int, name, gender string) *Immortal {
	if age < 0 {
		return nil
	}
	return &Immortal{Name: name, Gender: name, Age: age}
}
