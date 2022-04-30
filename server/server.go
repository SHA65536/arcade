package server

import (
	"github.com/gliderlabs/ssh"
)

type Server struct {
	Connection *ssh.Server
	Host       string
	Port       int
}
