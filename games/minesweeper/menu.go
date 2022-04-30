package minesweeper

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MineMenu struct {
	list   list.Model
	choice string
}

const listHeight = 14
const defaultWidth = 24

var (
	borderStyle        = lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("63"))
	titleStyle         = lipgloss.NewStyle().MarginLeft(3)
	itemStyle          = lipgloss.NewStyle().PaddingLeft(3)
	selectedArrowStyle = func(s string) string {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("63")).
			Blink(true).
			Bold(true).
			Render(s)
	}
	selectedItemStyle = func(s string) string {
		return selectedArrowStyle(">> ") +
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("#14F9D5")).
				Bold(true).
				Render(s)
	}
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 1 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle(s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

func (m *MineMenu) Init() tea.Cmd {
	return nil
}

func (m *MineMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter", " ":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *MineMenu) View() string {
	return "\n" + borderStyle.Render(m.list.View())
}

func NewMineMenu() *MineMenu {
	items := []list.Item{
		item("Easy   - 10x10 10 Mines  "),
		item("Medium - 16x16 40 Mines  "),
		item("Hard   - 16x30 99 Mines  "),
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose Difficulty"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowFilter(false)
	l.SetShowPagination(false)
	l.SetShowHelp(false)
	return &MineMenu{list: l}
}
