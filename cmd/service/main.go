package main

import (
	"context"
	"log/slog"

	"gitlab16.skiftrade.kz/libs-go/logger"
	application "gitlab16.skiftrade.kz/libs-go/new-app"
	"gitlab16.skiftrade.kz/templates/go/internal/repository"
	repoModels "gitlab16.skiftrade.kz/templates/go/internal/repository/models"
	"gitlab16.skiftrade.kz/templates/go/internal/service"
	"gitlab16.skiftrade.kz/templates/go/internal/usecase"
	api "gitlab16.skiftrade.kz/templates/go/pkg/api"
	"gitlab16.skiftrade.kz/templates/go/pkg/config"
)

func main() {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	cfg := application.LoadConfigFromEnv()

	app, err := application.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	addConfig := config.LoadConfig()

	err = app.Init(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to init app", logger.ErrorAttr(err))
		return
	}

	pool, err := repository.NewPostgres(ctx, repoModels.ConfigPostgres(addConfig.Postgres))
	if err != nil {
		slog.ErrorContext(ctx, "failed to init postgres", logger.ErrorAttr(err))
		return
	}
	defer pool.Close()
	repo := repository.NewRepository(pool)

	api.RegisterUsersServer(app.GrpcServer, service.NewService(usecase.NewUseCase(repo)))
	err = api.RegisterUsersHandler(ctx, app.ServeMux, app.GrpcConn)
	if err != nil {
		slog.ErrorContext(ctx, "failed to register users handler", logger.ErrorAttr(err))
		return
	}

	err = app.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to run app", logger.ErrorAttr(err))
		return
	}
}
