package minesweeper

import (
	"github.com/SHA65536/arcade/games"
	"github.com/SHA65536/arcade/session"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	mineMenu = iota
	mineStarting
	mineStarted
	mineWin
	mineLoss
	mineQuit
)

type Minsweeper struct {
	Session *session.Session
	State   int
	Menu    *MineMenu
	Board   *Board
}

func MakeMinesweeper(s *session.Session) games.Game {
	return &Minsweeper{
		Session: s,
		State:   mineMenu,
		Menu:    NewMineMenu(),
	}
}

func (m *Minsweeper) Init() tea.Cmd {
	return nil
}

func (m *Minsweeper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.State {
	case mineMenu:
		_, cmd = m.Menu.Update(msg)
		if m.Menu.choice != -1 {
			m.State = mineStarting
			m.Board = NewBoard(m.Menu.choice)
		}
	}
	return m, cmd
}

func (m *Minsweeper) View() string {
	switch m.State {
	case mineMenu:
		return m.Menu.View()
	case mineStarting:
		return m.Board.View()
	}
	return "Error"
}

func (m *Minsweeper) Status() int {
	if m.State == mineQuit {
		return games.FinishedStatus
	}
	return games.ActiveStatus
}

func (m *Minsweeper) Redirect() string {
	return "Menu"
}
