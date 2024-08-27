// gi/gi.go
package gi

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var mapMap [][]int

func mapGenerator(width, height int) {
	mapMap = make([][]int, height)
	for i := range mapMap {
		mapMap[i] = make([]int, width)
	}
	for i := range mapMap {
		for j := range mapMap[i] {
			if i == j {
				mapMap[i][j] = 1
			} else {
				mapMap[i][j] = 0
			}
		}
	}
}

func createMapContent(width, height int) string {
	mapGenerator(width, height)
	content := ""
	for i := range mapMap {
		for j := range mapMap[i] {
			switch mapMap[i][j] {
			case 0:
				content += " "
			case 1:
				content += "#"
			default:
				content += " "
			}
		}
		content += "\n"
	}
	return content
}

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
	mapContent := mapFrameStyle.Render(createMapContent(5, 5))
	statContent := statFrameStyle.Render("Str: 10 \nInt: 10")

	return fmt.Sprintf("%s\n%s%s%s%s", clearScreen, mapTitle, mapContent, statTitle, statContent)
}
