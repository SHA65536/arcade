package assets

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	MenuHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#14F9D5")).
			Background(lipgloss.Color("233")).
			Underline(true).
			Bold(true).
			Width(28).
			Align(lipgloss.Center).
			Render
	MenuItem = lipgloss.NewStyle().
			Background(lipgloss.Color("233")).
			Width(28).
			Align(lipgloss.Center).Render
	MenuArrow = func(s string) string {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("233")).
			Foreground(lipgloss.Color("63")).
			Blink(true).
			Bold(true).
			Render(s)
	}
	MenuSelect = func(s string) string {
		return MenuItem(MenuArrow(">>") +
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("#14F9D5")).
				Background(lipgloss.Color("233")).
				Bold(true).
				Render(s) +
			MenuArrow("<<"))
	}
)
