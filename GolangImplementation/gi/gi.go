// gi/gi.go
package gi

import (
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
	frameWidth := m.Width - 2
	frameHeight := (m.Height * 2) / 3

	// Создаём стиль для таба
	tabStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#0000ff")).
		Padding(0, 1).
		MarginBottom(-1) // Поднятие таба, чтобы он выглядел прикрепленным

	// Создаём стиль для рамки
	frameStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(frameWidth).
		Height(frameHeight)

	// Заголовок таба
	tab := tabStyle.Render("Map")

	// Содержимое рамки
	content := "Map goes here"
	frame := frameStyle.Render(content)

	// Объединяем таб и рамку
	return lipgloss.JoinVertical(lipgloss.Left, tab, frame)
}
