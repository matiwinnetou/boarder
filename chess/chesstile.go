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

func NewTile(b *core.Board, figure int) *ChessTile {
	t := &core.Tile{Board: b, Figure: figure}

	return &ChessTile{Tile: t}
}
