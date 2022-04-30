package server

import (
	"github.com/SHA65536/arcade/games"
	"github.com/SHA65536/arcade/games/minesweeper"
	"github.com/SHA65536/arcade/session"
)

var ArcadeMap = map[string]func(s *session.Session) games.Game{
	"Menu":        MakeMenu,
	"Minesweeper": minesweeper.MakeMinesweeper,
}
