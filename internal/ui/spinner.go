package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SpinnerModel struct {
	err error
	spinner.Model
	quitting bool
}

func NewSpinner() SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return SpinnerModel{Model: s}
}

func (s SpinnerModel) Init() tea.Cmd {
	return s.Tick
}

func (s SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		s.err = msg
		return s, nil

	default:
		var cmd tea.Cmd
		s.Model, cmd = s.Model.Update(msg)
		return s, cmd
	}
}

func (s SpinnerModel) View() string {
	if s.err != nil {
		return s.err.Error()
	}
	str := fmt.Sprintf("\n%s\n", s.Model.View())
	if s.quitting {
		return str + "\n"
	}
	return str
}
