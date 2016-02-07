package chess

import (
	"github.com/matiwinnetou/boarder/core"
)

type ChessBoard struct {
	*core.Board
}

func NewChessBoard() *ChessBoard {
	h := 8
	w := 8
	tiles := make([][]*core.Tile, 0)

	board := &core.Board{Tiles: tiles, Height: h, Width: w}
	board.Tiles = initEmptyTiles(h, w)

	return &ChessBoard{board}
}

func NewChessBoardWithTable(ct *ChessTable) *ChessBoard {
	board := NewChessBoard()
	board.Table = ct.Table

	return board
}

func initEmptyTiles(h, w int) [][]*core.Tile {
	tiles := make([][]*core.Tile, h)

	for i := range tiles {
		row := make([]*core.Tile, w)
		row[i] = NewChessTile(core.EMPTY_TILE).Tile
		tiles = append(tiles, row)
	}

	return tiles
}
