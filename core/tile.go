package core

const (
	EMPTY_TILE = 0
)

type Tile struct {
	Figure int
}

func (ct *Tile) IsEmpty() bool {
	return ct.Figure == 0
}
