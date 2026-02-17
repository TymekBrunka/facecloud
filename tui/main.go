package main

// A simple program that counts down from 5 and then exits.

import (
	"log"
	"os"
	"fctui/pages"

	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	pages.Current = pages.NewMainPage();

	// Initialize our program
	p := tea.NewProgram(pages.PageResolver)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
