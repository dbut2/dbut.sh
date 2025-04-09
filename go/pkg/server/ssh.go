package server

import (
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"

	"dbut.sh/pkg/provider"
	"dbut.sh/pkg/utils"
)

type SSHServer struct {
	Address     string
	HostKeyPath string
	provider    *provider.MarkdownProvider
	server      *ssh.Server
}

func NewSSHServer(address string, hostKeyPath string) (*SSHServer, error) {
	md, err := provider.NewMarkdownProvider(80)
	if err != nil {
		return nil, err
	}

	return &SSHServer{
		Address:     address,
		HostKeyPath: hostKeyPath,
		provider:    md,
	}, nil
}

func (s *SSHServer) Start() error {
	hostKey := utils.Must(os.ReadFile(s.HostKeyPath))

	server := utils.Must(wish.NewServer(
		wish.WithAddress(s.Address),
		wish.WithHostKeyPEM(hostKey),
		wish.WithMiddleware(
			lm.Middleware(),
			bm.Middleware(func(session ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, _ := session.Pty()
				m := s.newModel(pty.Window.Width, pty.Window.Height)
				return m, []tea.ProgramOption{tea.WithAltScreen()}
			}),
		),
	))

	s.server = server
	return server.ListenAndServe()
}

type model struct {
	sshServer *SSHServer
	viewport  viewport.Model
	width     int
	height    int
}

func (s *SSHServer) newModel(width, height int) *model {
	vp := viewport.New(width, height)

	m := &model{
		sshServer: s,
		viewport:  vp,
		width:     width,
		height:    height,
	}

	m.loadContent()
	return m
}

func (m *model) loadContent() {
	p, err := provider.NewMarkdownProvider(m.width)
	if err != nil {
		return
	}

	content, err := p.GetContent("index")
	if err != nil {
		return
	}

	m.viewport.SetContent(string(content))
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
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height

		m.loadContent()
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return m.viewport.View()
}

func StartSSHServer(address string, hostKeyPath string) error {
	s, err := NewSSHServer(address, hostKeyPath)
	if err != nil {
		return err
	}
	return s.Start()
}
