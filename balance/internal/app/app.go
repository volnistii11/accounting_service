package app

import (
	"context"
	"log"
	"sync"

	api "github.com/volnistii11/accounting_service/balance/internal/app/api/grpc"
	"github.com/volnistii11/accounting_service/balance/internal/app/server/grpc"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	conf := newConfig(cliFlags)

	ctx, _ := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	server := grpc.NewServer(conf.serverPort, api.NewHandler())
	if err := server.ListenAndServe(ctx, wg); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}

	return nil
}
