package selectionlist

import (
	"github.com/charmbracelet/bubbles/list"
	// tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Name, Desc string
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Name }

func New(items []list.Item) list.Model {
	return list.New(items, list.NewDefaultDelegate(), 0, 0);
}
