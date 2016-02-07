package chess

import (
	"github.com/matiwinnetou/boarder/core"
)

type ChessBoard struct {
	*core.Board
}

func NewChessBoard(ct *ChessTable) *ChessBoard {
	l := 8
	w := 8
	tiles := make([][]*core.Tile, 0)

	board := &core.Board{Table: ct.Table, Tiles: tiles, Length: l, Width: w}
    //board.Tiles = initEmptyTiles(board, l, w)

	return &ChessBoard{board}
}

//func initEmptyTiles(b *core.Board, l, w int) []*core.Tile {
//	tiles := make([]*core.Tile, l * w)
//
//	for i := 0; i < l; i++ {
//		for j := 0; j< w; j++ {
//			t := NewTile(b, core.EMPTY_TILE)
//			tiles[i][j] = t
//		}
//	}
//
//	return tiles
//}
