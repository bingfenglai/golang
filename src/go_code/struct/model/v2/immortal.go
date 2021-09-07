package v2

type immortal struct {
	name   string
	age    int
	gender string
}

func NewImmortal(age int, name, gender string) *immortal {
	if age < 0 {
		return nil
	}

	return &immortal{
		name:   name,
		age:    age,
		gender: gender,
	}
}
