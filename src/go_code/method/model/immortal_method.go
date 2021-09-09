package model


func NewImmortal(age int, name, gender string) *immortal {

	return &immortal{
		name:   name,
		age:    age,
		gender: gender,
	}
}


func (recv *immortal) GetName() string {
	return recv.name
}


