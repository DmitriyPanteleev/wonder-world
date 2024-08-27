// gi/gi.go
package gi

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width  int
	Height int
}

func NewProgram() *tea.Program {
	return tea.NewProgram(Model{})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	clearScreen := "\033[H\033[2J"

	frameWidth := m.Width - 2
	mapFrameHeight := (m.Height * 2) / 3
	statFrameHeight := m.Height - mapFrameHeight - 8

	mapFrameStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(frameWidth).
		Height(mapFrameHeight)

	statFrameStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(frameWidth).
		Height(statFrameHeight)

	titleStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(frameWidth).
		Align(lipgloss.Left)

	mapTitle := titleStyle.Render("Map")
	statTitle := titleStyle.Render("Stat")
	mapContent := mapFrameStyle.Render("Map goes here")
	statContent := statFrameStyle.Render("Str: 10 \nInt: 10")

	return fmt.Sprintf("%s\n%s%s%s%s", clearScreen, mapTitle, mapContent, statTitle, statContent)
}
