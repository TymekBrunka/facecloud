package pages

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type item_t struct {
	Name, Desc  string
	Key         int16
	Hidden      bool
	PreHandler  func(*EnvEditorMain, *item_t) error
	PostHandler func(*EnvEditorMain, *item_t) error
}

func (i item_t) Title() string       { return i.Name }
func (i item_t) Description() string { return i.Desc }
func (i item_t) FilterValue() string { return i.Name }

func newList(items []list.Item) list.Model {
	return list.New(items, list.NewDefaultDelegate(), 0, 0)
}

type EnvEditorMain struct {
	parent Page
	list   list.Model
	text   textinput.Model

	selectedItemIndex int16
	givesInput        bool
	width             int
	height            int
}

func (p EnvEditorMain) Init() tea.Cmd {
	return nil
}

func (p EnvEditorMain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "escape":
			if p.givesInput {
				p.givesInput = false
				p.selectedItemIndex = -1
				p.text.Blur()
				return p, tea.WindowSize()
			}
			Current = p.parent
			return Current, tea.WindowSize()
		case "enter":
			if p.givesInput {
				item, _ := ENVEDITORMAIN_list[p.selectedItemIndex].(item_t)
				if item.PostHandler != nil {
					if err := item.PostHandler(&p, &item); err != nil {
						log.Println("Error saving to a field", item.Key, err)
						return p, tea.WindowSize()
					}
				} else {
					Env[keys[item.Key]] = p.text.Value()
				}

				//save to .env
				for ok := true; ok; ok = true {
					envstring, err := godotenv.Marshal(Env)
					if err != nil {
						log.Println("Error producting .env file", err)
						break
					}

					f, err := os.Create(".env")
					if err != nil {
						log.Println("Error creating .env file", err)
						break
					}

					bytessaved, err := f.WriteString(envstring)
					if err != nil {
						log.Println("Error saving .env file", err)
						break
					}

					log.Println("Saved", bytessaved, "bytes to .env file")
					break
				}

				p.givesInput = false
				p.selectedItemIndex = -1
				p.text.Blur()
				return p, tea.WindowSize()
			} else {
				p.givesInput = true
				p.selectedItemIndex = int16(p.list.Index())
				item, _ := ENVEDITORMAIN_list[p.selectedItemIndex].(item_t)
				p.text.SetValue(Env[keys[item.Key]])
				p.text.Focus()
				if item.PreHandler != nil {
					if err := item.PreHandler(&p, &item); err != nil {
						log.Println("Error opening a field", item.Key, err)
						return p, tea.WindowSize()
					}
				}
				return p, tea.WindowSize()
			}
		}
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		var padding int = 0
		if p.givesInput {
			padding = 4
		}
		p.list.SetSize(msg.Width, msg.Height-padding)
		return p, nil
	}

	var cmd tea.Cmd
	if p.givesInput {
		p.text, cmd = p.text.Update(msg)
	} else {
		p.list, cmd = p.list.Update(msg)
	}
	return p, cmd
}

func (p EnvEditorMain) View() string {
	var additional string
	if p.givesInput {
		item, _ := ENVEDITORMAIN_list[p.selectedItemIndex].(item_t)
		additional += fmt.Sprintf("\n--------\nEdytowanie %s\n -> ", keys[item.Key])
		if item.Hidden {
			additional += strings.Repeat("*", len(p.text.Value()))
		} else {
			additional += p.text.Value()
		}
	}
	return p.list.View() + additional
}

func (p EnvEditorMain) GetParent() Page {
	return p.parent
}

func (p EnvEditorMain) DeletePage() {
	Current = p.parent
}

func password_pre_handler(p *EnvEditorMain, item *item_t) error {
	p.text.SetValue("")
	return nil
}

func password_post_handler(p *EnvEditorMain, item *item_t) error {
	h := sha256.New()
	password := p.text.Value()
	if len(password) < 5 {
		return fmt.Errorf("hasło powinno mieć co najmniej 5 znaków")
	}
	h.Write([]byte(password[4:5] + password + password[2:4]))
	Env[keys[item.Key]] = hex.EncodeToString(h.Sum(nil))
	return nil
}

var ENVEDITORMAIN_list = []list.Item{
	item_t{Key: DB, Name: "BAZA DANYCH", Desc: "klucz do bazy danych"},
	item_t{Key: REINIT_LOGIN, Name: "login resetowania", Desc: "login (nazwa użytkownika) wykorzystywane do funkcji resetowania bazy"},
	item_t{Key: REINIT_PASSWORD,
		PreHandler: password_pre_handler, PostHandler: password_post_handler,
		Hidden: true, Name: "hasło resetowania", Desc: "hasło wykorzystywane do funkcji resetowania bazy"},
	item_t{Key: SUPERUSER_EMAIL, Name: "email superużytkownika", Desc: "(zostanie zastosowany po zresetowaniu bazy z facecloudem)"},
	item_t{Key: SUPERUSER_PASSWORD,
		PreHandler: password_pre_handler, PostHandler: password_post_handler,
		Hidden: true, Name: "hasło superużytkownika", Desc: "(zostanie zastosowane po zresetowaniu bazy z facecloudem)"},
	item_t{Key: SUPERUSER_BIRTH_DATE, Name: "data urodzenia superużytkownika", Desc: "(zostanie zastosowana po zresetowaniu bazy z facecloudem)"},
	item_t{Name: "-- DANE DO TESTOWANIA --", Desc: "vvvvvvvvvvvvvvvvvv"},
	item_t{Key: TEST_DB, Name: "TESTOWA BAZA DANYCH", Desc: "klucz do bazy danych do testów jednostkowych"},
	item_t{Key: TEST_REINIT_LOGIN, Name: "testowy login resetowania", Desc: "login (nazwa użytkownika) wykorzystywane do funkcji resetowania bazy do testów jednostkowych"},
	item_t{Key: TEST_REINIT_PASSWORD,
		PreHandler: password_pre_handler, PostHandler: password_post_handler,
		Hidden: true, Name: "testowe hasło resetowania", Desc: "hasło wykorzystywane do funkcji resetowania bazy do testów jednostkowych"},
	item_t{Key: TEST_SUPERUSER_EMAIL, Name: "testowy email superużytkownika", Desc: "email superużytkownika do testów jednostkowych"},
	item_t{Key: TEST_SUPERUSER_PASSWORD,
		PreHandler: password_pre_handler, PostHandler: password_post_handler,
		Hidden: true, Name: "testowe hasło superużytkownika", Desc: "hasło superużytkownika do testów jednostkowych"},
	item_t{Key: TEST_SUPERUSER_BIRTH_DATE, Name: "testowa data urodzenia superużytkownika", Desc: "data urodzenia superużytkownika do testów jednostkowych"},
}

func NewEnvEditorMainPage() EnvEditorMain {
	list := newList(ENVEDITORMAIN_list)
	text := textinput.New()

	p := EnvEditorMain{
		parent: nil,
		list:   list,
		text:   text,
		width:  0,
		height: 0,
	}

	return p
}
