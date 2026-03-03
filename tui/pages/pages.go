package pages

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
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

const (
	DB = iota
	REINIT_LOGIN
	REINIT_PASSWORD
	SUPERUSER_EMAIL
	SUPERUSER_PASSWORD
	SUPERUSER_BIRTH_DATE

	TEST_DB
	TEST_REINIT_LOGIN
	TEST_REINIT_PASSWORD
	TEST_SUPERUSER_EMAIL
	TEST_SUPERUSER_PASSWORD
	TEST_SUPERUSER_BIRTH_DATE
)

var keys []string = []string{
	"DB",
	"REINIT_LOGIN",
	"REINIT_PASSWORD",
	"SUPERUSER_EMAIL",
	"SUPERUSER_PASSWORD",
	"SUPERUSER_BIRTH_DATE",

	"TEST_DB",
	"TEST_REINIT_LOGIN",
	"TEST_REINIT_PASSWORD",
	"TEST_SUPERUSER_EMAIL",
	"TEST_SUPERUSER_PASSWORD",
	"TEST_SUPERUSER_BIRTH_DATE",
}

var Env map[string]string

func (p PageResolver_t) Init() tea.Cmd {
	return nil
}

func (p PageResolver_t) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return Current.Update(msg)
}

func (p PageResolver_t) View() string {
	return Current.View()
}

func Run() {

	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	Current = NewMainPage()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	if Env, err = godotenv.Read(); err != nil {
		log.Println("Error reading the environment variables: %v", err)
	}

	// log.Printf("%+v\n", Env)

	p := tea.NewProgram(PageResolver)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
