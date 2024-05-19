package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	desc "github.com/volnistii11/accounting_service/balance/pkg/api/transfer/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	port     string
	handlers desc.TransferServer
}

func NewServer(port string, handlers desc.TransferServer) *Server {
	return &Server{
		port:     port,
		handlers: handlers,
	}
}

func (s *Server) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)
	desc.RegisterTransferServer(server, s.handlers)

	stopCh := make(chan struct{})
	errCh := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(stopCh)
		log.Printf("grpc server listening at %v", lis.Addr())
		if err = server.Serve(lis); err != nil {
			errCh <- err
			return
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("shutting down gRPC server by ctx")
		server.GracefulStop()
		return nil
	case <-stopCh:
		log.Println("grpc server stopped")
	case err = <-errCh:
		return err
	}

	return nil
}
