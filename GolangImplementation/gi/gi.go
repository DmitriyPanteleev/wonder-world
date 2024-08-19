// gi/gi.go
package gi

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
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
	frameWidth := m.Width - 1
	frameHeight := (m.Height * 2) / 3

	frameStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(frameWidth).
		Height(frameHeight)

	contentStyle := lipgloss.NewStyle().
		MarginTop(1).
		MarginLeft(2)

	titleStyle := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Width(frameWidth).
		Render("Map")

	content := "Здесь может быть ваш контент"

	return titleStyle + "\n" + frameStyle.Render(contentStyle.Render(content))
}
