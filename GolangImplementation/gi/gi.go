// gi/gi.go
package gi

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var mapMap [][]int
var xPos, yPos int
var backupCell int

func mapGenerator(width, height int) {
	if len(mapMap) != 0 {
		return // Map has already been generated
	}

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
	xPos = rand.Intn(width)
	yPos = rand.Intn(height)
	mapMap[yPos][xPos] = 8
}

func createMapContent() string {
	content := ""
	for i := range mapMap {
		for j := range mapMap[i] {
			switch mapMap[i][j] {
			case 0:
				content += "."
			case 1:
				content += "#"
			case 8:
				content += "@"
			default:
				content += "."
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
	// Generate the map once during initialization
	mapGenerator(10, 10)
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
		case "up":
			if yPos > 0 {
				mapMap[yPos][xPos] = backupCell
				yPos--
				backupCell = mapMap[yPos][xPos]
				mapMap[yPos][xPos] = 8
			}
		case "down":
			if yPos < len(mapMap)-1 {
				mapMap[yPos][xPos] = backupCell
				yPos++
				backupCell = mapMap[yPos][xPos]
				mapMap[yPos][xPos] = 8
			}
		case "left":
			if xPos > 0 {
				mapMap[yPos][xPos] = backupCell
				xPos--
				backupCell = mapMap[yPos][xPos]
				mapMap[yPos][xPos] = 8
			}
		case "right":
			if xPos < len(mapMap[0])-1 {
				mapMap[yPos][xPos] = backupCell
				xPos++
				backupCell = mapMap[yPos][xPos]
				mapMap[yPos][xPos] = 8
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
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
	mapContent := mapFrameStyle.Render(createMapContent())
	statContent := statFrameStyle.Render("Str: 10 \nInt: 10")

	return fmt.Sprintf("%s%s%s%s", mapTitle, mapContent, statTitle, statContent)
}
