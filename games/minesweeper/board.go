package minesweeper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SHA65536/arcade/assets"
	tea "github.com/charmbracelet/bubbletea"
)

type Board struct {
	Tiles [][]Tile
	Start int64
	Mode  string
	SizeX int
	SizeY int
	Mines int
	CurX  int
	CurY  int
}

type Tile struct {
	Value  int
	Opened bool
	Flag   bool
}

var (
	aroundX = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	aroundY = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
)

func NewBoard(mode int) *Board {
	board := &Board{SizeX: 10, SizeY: 10, Mines: 10}
	if mode == 1 {
		board.SizeX = 16
		board.SizeY = 16
		board.Mines = 40
	} else if mode == 2 {
		board.SizeX = 16
		board.SizeY = 30
		board.Mines = 99
	}
	board.GenTiles()
	return board
}

func (m *Board) GenTiles() {
	seed := rand.NewSource(time.Now().UnixNano())
	rGen := rand.New(seed)
	// Making board
	m.Tiles = make([][]Tile, m.SizeY)
	for row := range m.Tiles {
		m.Tiles[row] = make([]Tile, m.SizeX)
	}
	// Making mines
	for i := 0; i < m.Mines; i++ {
		mX, mY := rGen.Intn(m.SizeX), rGen.Intn(m.SizeY)
		if m.Tiles[mY][mX].Value == -1 {
			i--
		} else {
			m.Tiles[mY][mX].Value = -1
		}
	}
	// Making numbers
	for y := 0; y < m.SizeY; y++ {
		for x := 0; x < m.SizeX; x++ {
			var sum int
			if m.Tiles[y][x].Value == -1 {
				continue
			}
			for i := range aroundX {
				var nY, nX = y + aroundY[i], x + aroundX[i]
				if nY >= m.SizeY || nY < 0 || nX >= m.SizeX || nX < 0 {
					continue
				}
				if m.Tiles[nY][nX].Value == -1 {
					sum++
				}
			}
			m.Tiles[y][x].Value = sum
		}
	}
}

func (m *Board) Init() tea.Cmd {
	return nil
}

func (m *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m *Board) View() string {
	flat := flatten(m.Tiles)
	return fmt.Sprintf(assets.EasyBoard, flat...)
}

func flatten(slices [][]Tile) []any {
	var total int
	for _, s := range slices {
		total += len(s)
	}
	flat := make([]any, 0, total)
	for _, s := range slices {
		for _, v := range s {
			if v.Value == -1 {
				flat = append(flat, "x")
			} else {
				flat = append(flat, fmt.Sprint(v.Value))
			}
		}
	}
	return flat
}
