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
	//board.Tiles = initEmptyTiles(board)
	//placePawns(board, board.Tiles)

	chessBoard := &ChessBoard{board}
	board.Tiles = initEmptyTiles(chessBoard)

	return chessBoard;
}

func initEmptyTiles(b *ChessBoard) [][]*ChessTile {
	tiles := make([][]*ChessTile, b.Height)

	for i := range tiles {
		row := make([]*ChessTile, b.Width)
		row[i] = NewChessTile(core.EMPTY_TILE, NO_COLOR)
		tiles = append(tiles, row)
	}

	return tiles
}

//func placePawns(b *core.Board, tiles [][]*ChessTile)  {
//	for i := 0; i <= b.Width; i++ {
//		tiles[1][i] = NewChessTile(PAWN_TILE, WHITE_COLOR)
//		tiles[b.Height - 2][i] = NewChessTile(PAWN_TILE, BLACK_COLOR)
//	}
//}
