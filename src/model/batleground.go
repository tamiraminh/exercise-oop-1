package model

type Batleground struct {
	surface  [][]string
	gridSize int
}

func NewBattleground(size int) *Batleground {
	b := Batleground{}
	b.gridSize = size

	for i := 0; i < size; i++ {
		row := make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = "-"
		}
		b.surface = append(b.surface, row)

	}

	return &b
}

func (b Batleground) IsShip(x, y int) bool {
	return b.surface[x][y] == "S"
}

func (b Batleground) IsEmpty(x, y int) bool {
	return b.surface[x][y] == "-"
}

func (b Batleground) SetShip(x, y int) {
	b.surface[x][y] = "B"
}

func (b Batleground) SetMissile(x, y int) {
	b.surface[x][y] = "O"
}

func (b Batleground) SetDeadShip(x, y int) {
	b.surface[x][y] = "X"
}
