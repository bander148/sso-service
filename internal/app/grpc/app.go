package grpc

import (
	"log/slog"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCSserver := grpc.NewServer()

	authgrpc.Register(gRPCSserver)
	return &App{
		log:        log,
		gRPCServer: gRPCSserver,
		port:       port,
	}
}
