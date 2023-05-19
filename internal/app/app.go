package app

import (
	"github.com/homaderaka/peersmsg"
	"pnode/internal/server"
)

type App struct {
	s *server.Server
}

func New() (app *App, err error) {
	p := peersmsg.NewParser('\x00')

	serv := server.NewServer(p)

	app = &App{s: serv}
	return
}

func (app *App) Run() {
	app.s.AcceptConnections()
}
