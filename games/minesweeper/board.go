package minesweeper

type Board struct {
	Tiles [][]Tile
	Mode  string
	Size  int
	Mines int
	CurX  int
	CurY  int
}

type Tile struct {
	Value  int
	Opened bool
}
