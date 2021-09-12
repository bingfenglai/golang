package model

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 凡人
type Mortal struct {
	name ,
	gender ,
	spiritualRoot string
	age int

}

func NewMortal(name, gender string, age int) *Mortal {
	mortal:=&Mortal{
		name:   name,
		gender: gender,
		age:    age,
	}
	mortal.GenSpiritualRootNames()
	return mortal
}

func (recv *Mortal) SpiritualRoot() string {
	if &recv.spiritualRoot == nil{
		return "没有灵根"
	}

	return recv.spiritualRoot
}

func (recv *Mortal)Practice()  {
	fmt.Println(recv.name,"开始修行...")
}

func (recv *Mortal) GenSpiritualRootNames() {
	gsrn := []string{
		"金灵根","水灵根","木灵根","火灵根","土灵根","没有灵根",
	}
	index, _ := rand.Int(rand.Reader, big.NewInt(5))

	recv.spiritualRoot =  gsrn[index.Int64()]
}
