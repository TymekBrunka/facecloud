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
		case "enter":
			if p.list.SelectedItem() == MAINPAGE_list[0] {
				page := NewEnvEditorMainPage()
				page.parent = p
				Current = page
				return Current, tea.WindowSize() // the size is queried to fix broken rendering after loading new page
			}
			return Current, tea.WindowSize()
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

var MAINPAGE_list = []list.Item{
	sl.Item{Name: "Konfiguracja środowiska", Desc: "Konfigurator pliku .env (zmiany zachodzą po ponownym uruchomieniu serwera)"},
}

func NewMainPage() MainPage {
	list := sl.New(MAINPAGE_list)
	p := MainPage{
		parent: nil,
		list:   list,
		width:  0,
		height: 0,
	}

	return p
}
