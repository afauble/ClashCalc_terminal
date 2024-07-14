package frontend

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func (m model) View() string {
	panic("unimplemented")
}

func initialModel() model {
	return model{
		choices:  []string{"Zap", "Earthquake", "Fireball", "Giantarrow"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Clash Calc")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:

		// Check which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing
	return m, nil
}
