package app

import (
	"context"
	"github.com/homaderaka/peersmsg"
	"pnode/internal/server"
	"pnode/internal/service"
	"pnode/internal/storage"
)

type App struct {
	s *server.Server
}

func New() (app *App, err error) {
	p := peersmsg.NewParser('\x00')

	s := storage.NewStorageRAM()

	storageService := service.NewService(s)

	serv := server.NewServer(p, storageService)

	app = &App{s: serv}
	return
}

func (app *App) Run(c context.Context) {
	app.s.AcceptConnections(c)
}
