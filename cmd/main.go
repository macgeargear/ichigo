package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/macgeargear/ichigo/internal/component"
)

func main() {
	columns := []table.Column{
		{Title: "Name", Width: 4},
		{Title: "Meaning", Width: 10},
		{Title: "Reading", Width: 10},
		{Title: "Example", Width: 10},
	}

	rows := []table.Row{
		{"九", "nine", "く", "九つ"},
		{"十", "ten", "じゅう", "十日"},
		{"百", "hundred", "ひゃく", "百円"},
		{"千", "thousand", "せん", "千円"},
		{"万", "ten thousand", "まん", "万年"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := component.TableModel{
		Table: t,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
