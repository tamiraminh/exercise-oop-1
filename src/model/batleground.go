package model

type Batleground struct {
	surface  [][]string
	gridSize int
}

func NewBattleground(size int) *Batleground {
	b := Batleground{}
	return &b
}
