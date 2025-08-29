package main

import (
	"context"
	repoModels "gitlab16.skiftrade.kz/templates1/go/internal/repository/models"
	"log/slog"

	"gitlab16.skiftrade.kz/libs-go/logger"
	application "gitlab16.skiftrade.kz/libs-go/new-app"
	"gitlab16.skiftrade.kz/templates1/go/internal/repository"
	"gitlab16.skiftrade.kz/templates1/go/internal/service"
	"gitlab16.skiftrade.kz/templates1/go/internal/usecase"
	api "gitlab16.skiftrade.kz/templates1/go/pkg/api"
)

func main() {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	cfg := application.LoadConfigFromEnv()

	app, err := application.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	err = app.Init(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to init app", logger.ErrorAttr(err))
		return
	}

	pool, err := repository.NewPostgres(ctx, repoModels.ConfigPostgres{})
	if err != nil {
		return
	}

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

	// Закрываем соединение с БД
	repo.Close()
}
