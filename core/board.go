package core

type Board struct {
	Table  *Table
	Tiles  [][]*Tile
	Height int
	Width  int
}

func (b *Board) TileAt(x, y int) *Tile {
	return b.Tiles[x][y]
}

func (b *Board) PlaceTileAt(t *Tile, x, y int) {
	b.Tiles[x][y] = t
}
