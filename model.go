package main

import (
	"image/color"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/luism2302/chess-server/internal/game"
)

type model struct {
	game      *game.Game
	width     int
	height    int
	moveField textinput.Model
	styles    *styles
}

func newModel() *model {
	mf := textinput.New()
	mf.SetWidth(30)
	mf.Placeholder = "Your move here: e.g e2-e4"
	mf.Focus()
	return &model{
		game:      game.New(),
		moveField: mf,
		styles:    defaultStyles(),
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			move := m.moveField.Value()
			err := m.game.Update(move)
			if err != nil {
				return m, tea.Quit
			}
			m.moveField.Reset()
		}
	}
	m.moveField, cmd = m.moveField.Update(msg)
	return m, cmd
}

func (m *model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("loading...")
	}

	boardRepr := m.game.BoardRepr()
	s := lipgloss.JoinVertical(lipgloss.Center, boardRepr, m.styles.inputField.Render(m.moveField.View()))
	s = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
	view := tea.NewView(s)
	view.AltScreen = true

	return view
}

type styles struct {
	borderColor color.Color
	inputField  lipgloss.Style
}

func defaultStyles() *styles {
	s := new(styles)
	s.borderColor = lipgloss.BrightWhite
	s.inputField = lipgloss.NewStyle().BorderForeground(s.borderColor).Foreground(lipgloss.BrightWhite).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(40)
	return s
}
