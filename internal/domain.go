package internal

import (
	"context"
	rcModels "gitlab16.skiftrade.kz/templates1/go/internal/rest/client/models"

	"github.com/jackc/pgx/v5"
	repoModels "gitlab16.skiftrade.kz/templates1/go/internal/repository/models"
	ucModels "gitlab16.skiftrade.kz/templates1/go/internal/usecase/models"
)

type Repository interface {
	DBBeginTransaction(ctx context.Context) (pgx.Tx, error)

	ReadUser(ctx context.Context, id int64, dbTx pgx.Tx) (user repoModels.User, err error)
}

type UseCase interface {
	GetUser(ctx context.Context, input ucModels.GetUserInput) (output ucModels.GetUserOutput, err error)
}

type Client interface {
	PostingsToCancel(ctx context.Context, token string, req rcModels.PostingsToCancelReq) (rcModels.PostingsToCancelResp, error)
	PostingsCancelResponse(ctx context.Context, token string, req rcModels.PostingsCancelResponseReq) (rcModels.PostingsCancelResponseResp, error)
}
