package chess

import (
	"github.com/matiwinnetou/boarder/core"
)

const (
	PAWN_TILE = 1 << iota
	ROOK_TILE
	KING_TILE
	QUEEN_TILE
	KNIGHT_TILE
	BISHOP_TILE
)

type ChessTile struct {
	*core.Tile
}

func NewChessTile(figure int) *ChessTile {
	t := &core.Tile{Figure: figure}

	return &ChessTile{Tile: t}
}

func NewChessTileWithBoard(b *core.Board, figure int) *ChessTile {
	tile := NewChessTile(figure)
	tile.Board = b

	return tile
}
