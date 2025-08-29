package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
	repoModels "gitlab16.skiftrade.kz/templates1/go/internal/repository/models"
	ucModels "gitlab16.skiftrade.kz/templates1/go/internal/usecase/models"
)

type Repository interface {
	DBBeginTransaction(ctx context.Context) (pgx.Tx, error)
	Close()

	ReadUser(ctx context.Context, id int64, dbTx pgx.Tx) (user repoModels.User, err error)
}

type UseCase interface {
	GetUser(ctx context.Context, input ucModels.GetUserInput) (output ucModels.GetUserOutput, err error)
}
