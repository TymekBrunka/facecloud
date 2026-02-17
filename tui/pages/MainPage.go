package pages

import (
	sl "fctui/selectionList"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type MainPage struct {
	parent Page
	list   list.Model

	width  int
	height int
}

func (p MainPage) Init() tea.Cmd {
	return nil
}

func (p MainPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return p, tea.Quit
		}
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		p.list.SetSize(msg.Width, msg.Height)
		return p, nil
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p MainPage) View() string {
	return p.list.View()
}

func (p MainPage) GetParent() Page {
	return p.parent
}

func (p MainPage) DeletePage() {
	Current = p.parent
}

func NewMainPage() MainPage {
	list := sl.New([]list.Item{
		sl.Item{Name: "reinit password", Desc: "hasło do reinicjalizacji"},
		sl.Item{Name: "dyrek password", Desc: "hasło dyrka"},
	})

	p := MainPage{
		parent: nil,
		list:   list,
		width:  0,
		height: 0,
	}

	return p
}
