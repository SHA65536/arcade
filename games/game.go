package games

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ActiveStatus = iota
	FinishedStatus
)

type Game interface {
	tea.Model
	Status() int
	Redirect() string
}
