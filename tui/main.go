package main

// A simple program that counts down from 5 and then exits.

import (
	// "fmt"
	"log"
	"os"
	// "time"

	// "github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func CreateTextInput() textinput.Model {
	ti := textinput.New()
	ti.Focus()
	return ti
}

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	// Initialize our program
	p := tea.NewProgram(New())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func New() model {
	return model{
		0,
		0,
		CreateTextInput(),
		lipgloss.NewStyle().
						BorderForeground(lipgloss.Color("36")).
						BorderStyle(lipgloss.NormalBorder()).
						Padding(1).
						Width(80),
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
// type model int
type model struct {
	width  int
	height int
	input  textinput.Model
	style  lipgloss.Style
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}

	// case tickMsg:
	// 	m.i--
	// 	if m.i <= 0 {
	// 		return m, tea.Quit
	// 	}
	// 	return m, tick
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	// return fmt.Sprintf("Hi. This program will exit in %d seconds.\n\nTo quit sooner press ctrl-c, or press ctrl-z to suspend...\n", m.i)
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.style.Render(m.input.View()),
		))
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
// type tickMsg time.Time
//
// func tick() tea.Msg {
// 	time.Sleep(time.Second)
// 	return tickMsg{}
// }
