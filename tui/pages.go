package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page struct {
	Parent *Page
	Model tea.Model
}

var current Page

func (p Page) Delete() {
	current = *p.Parent;
}
