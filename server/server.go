package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	m := &SessionModel{
		Session: session.NewSession(),
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func (m *SessionModel) Init() tea.Cmd {
	return nil
}

func (m *SessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var retCmd tea.Cmd

	return m, retCmd
}

func (m *SessionModel) View() string {
	return m.Session.SessionId.String()
}
