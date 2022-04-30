package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SHA65536/arcade/games"
	"github.com/SHA65536/arcade/session"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

type Server struct {
	Connection *ssh.Server
	Host       string
	Port       int
}

type SessionModel struct {
	Session *session.Session
	Current games.Game
}

// MakeServer makes a new server object
func MakeServer(host string, port int) (*Server, error) {
	var err error
	serv := &Server{Host: host, Port: port}
	serv.Connection, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	return serv, err
}

// Run starts the server blockingly and gracefully shuts down on
// keyboard interrupt
func (srv *Server) Run() error {
	var err error
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting Arcade on %s:%d", srv.Host, srv.Port)
	go func() {
		if err = srv.Connection.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping Arcade")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := srv.Connection.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

// teaHandler Handles starting the bubbletea program
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}
	curSession := session.NewSession()
	m := &SessionModel{
		Session: curSession,
		Current: MakeMenu(curSession),
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func (m *SessionModel) Init() tea.Cmd {
	return nil
}

func (m *SessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var retCmd tea.Cmd

	// Checking for interrupt
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "ctrl+c" {
		return m, tea.Quit
	}

	// Updating current game
	if m.Current.Status() == games.ActiveStatus {
		_, retCmd = m.Current.Update(msg)
	}

	// Checking for redirect
	if m.Current.Status() == games.FinishedStatus {
		if newGame, ok := ArcadeMap[m.Current.Redirect()]; ok {
			// Showing starting screen of redirect
			m.Current = newGame(m.Session)
		}
	}
	return m, retCmd
}

func (m *SessionModel) View() string {
	switch m.Current.Status() {
	case games.ActiveStatus:
		return m.Current.View()
	case games.FinishedStatus:
		return fmt.Sprintf("Invalid Redirect Index: %s", m.Current.Redirect())
	}
	return "Error, Invalid Status"
}
