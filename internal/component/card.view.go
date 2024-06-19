package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	hotPink  = lipgloss.Color("#ff06b7")
	darkGray = lipgloss.Color("#767676")
)

const (
	name = iota
	reading
	meaning
)

type Card struct {
	inputs []textinput.Model

	height int
	width  int
	style  lipgloss.Style
}

func initCard() Card {
	inputs := make([]textinput.Model, 3)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Name"

	return Card{
		inputs: inputs,
		height: 5,
		width:  30,
		style:  lipgloss.NewStyle().Padding(1, 2),
	}
}

func (c Card) Init() tea.Cmd {
	return textinput.Blink
}

func (c Card) View() string {
	return ""
}
