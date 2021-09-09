package model

// 修士的灵根
type lingGen struct {
	linGenNames[] string
}

func NewLinggen(name ...string) *lingGen {
	return &lingGen{linGenNames: name}
}

func (recv *lingGen) LingGenNames() []string {
	return recv.linGenNames
}
