package pages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page interface {
	GetParent() Page
	DeletePage()
	tea.Model
}

var Current Page
var PageResolver PageResolver_t

type PageResolver_t struct {
	Width  int
	Height int
}

func (p PageResolver_t) Init() tea.Cmd {
	return nil
}

func (p PageResolver_t) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return Current.Update(msg)
}

func (p PageResolver_t) View() string {
	return Current.View()
}
