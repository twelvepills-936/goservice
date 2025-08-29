package usecase

import (
	"gitlab16.skiftrade.kz/templates1/go/internal"
)

type useCase struct {
	repo internal.Repository
}

func NewUseCase(
	repo internal.Repository) internal.UseCase {
	return &useCase{
		repo: repo,
	}
}
