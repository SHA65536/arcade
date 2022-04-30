package server

import (
	"fmt"

	"github.com/SHA65536/arcade/assets"
	"github.com/SHA65536/arcade/games"
	"github.com/SHA65536/arcade/session"
	tea "github.com/charmbracelet/bubbletea"
)

var mainMenu = []string{
	"Continue As Guest",
	"Login",
	"",
	"",
	"",
	"",
	"",
	"",
}

var gameMenu = []string{
	"TicTacToe",
	"Minesweeper",
	"",
	"",
	"",
	"",
	"",
	"",
}

type Menu struct {
	Session      *session.Session
	MenuIdx      int
	SelectionIdx int
	Finished     int
	Destination  string
}

func MakeMenu(s *session.Session) games.Game {
	menu := &Menu{
		Session: s,
	}
	if menu.Session.Username != "Guest" {
		menu.MenuIdx = 1
	}
	return menu
}

func (m *Menu) Init() tea.Cmd {
	return nil
}

func (m *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var menu *[]string
	if m.MenuIdx == 0 {
		menu = &mainMenu
	} else {
		menu = &gameMenu
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			if m.SelectionIdx > 0 {
				m.SelectionIdx -= 1
			}
		case "down", "s":
			if m.SelectionIdx < len(*menu)-2 {
				m.SelectionIdx += 1
			}
		case " ", "enter":
			if m.MenuIdx == 0 && m.SelectionIdx == 0 {
				m.MenuIdx = 1
			} else {
				m.Finished = games.FinishedStatus
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Menu) View() string {
	var menu *[]string
	if m.MenuIdx == 0 {
		menu = &mainMenu
	} else {
		menu = &gameMenu
	}
	rendered := []any{assets.MenuHeader("Welcome to the Arcade!"), assets.MenuHeader("")}
	for idx, item := range *menu {
		if idx == m.SelectionIdx {
			rendered = append(rendered, assets.MenuSelect(item))
		} else {
			rendered = append(rendered, assets.MenuItem(item))
		}
	}
	return fmt.Sprintf(assets.Gameboy, rendered...)
}

func (m *Menu) Status() int {
	return m.Finished
}

func (m *Menu) Redirect() string {
	if m.MenuIdx == 0 && m.SelectionIdx == 2 {
		return "Auth"
	}
	return gameMenu[m.SelectionIdx]
}
