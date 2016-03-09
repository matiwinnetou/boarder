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

	NO_COLOR
	BLACK_COLOR
	WHITE_COLOR
)

type ChessTile struct {
	*core.Tile
	Color int
}

func NewChessTile(figure, color int) *ChessTile {
	t := &core.Tile{Figure: figure}

	return &ChessTile{Tile: t, Color: color}
}
