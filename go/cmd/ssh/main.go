package main

import (
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"

	"dbut.sh/pages"
)

func main() {
	server, err := wish.NewServer(
		wish.WithAddress(":8022"),
		wish.WithHostKeyPEM(must(os.ReadFile("host_key"))),
		wish.WithMiddleware(
			lm.Middleware(),
			bm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				p, _, _ := s.Pty()
				m := newModel(p.Window.Width, p.Window.Height)
				return m, []tea.ProgramOption{tea.WithAltScreen()}
			}),
		),
	)
	if err != nil {
		panic(err.Error())
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

type model struct {
	viewport viewport.Model
}

func newModel(width, height int) *model {
	vp := viewport.New(width, height)
	m := &model{
		viewport: vp,
	}
	m.loadContent()
	return m
}

func (m *model) loadContent() {
	content, err := pages.MD.ReadFile("index.md")
	if err != nil {
		return
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.viewport.Width),
		glamour.WithEmoji(),
	)
	if err != nil {
		return
	}

	renderedContent, err := r.Render(string(content))
	if err != nil {
		return
	}

	m.viewport.SetContent(renderedContent)
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return m.viewport.View() + "\n\n" + "Hello, world!"
}

func must[t any](v t, err error) t {
	if err != nil {
		panic(err.Error())
	}
	return v
}
