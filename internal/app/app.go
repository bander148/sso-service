package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO : storage

	// TODO : init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPSrv: grpcApp,
	}
}
