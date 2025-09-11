package service

import (
	"context"
	"errors"

	"gitlab16.skiftrade.kz/templates/go/internal"
	ucModels "gitlab16.skiftrade.kz/templates/go/internal/usecase/models"
	api "gitlab16.skiftrade.kz/templates/go/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewService(uc internal.UseCase) *service {
	return &service{
		uc: uc,
	}
}

type service struct {
	api.UnimplementedUsersServer
	uc internal.UseCase
}

func (s *service) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	inp := ucModels.GetUserInput{UserID: req.GetUserId()}
	user, err := s.uc.GetUser(ctx, inp)
	if err != nil {
		switch {
		case errors.Is(err, ucModels.ErrUserIDIsInvalid):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, ucModels.ErrUserIsNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, ucModels.ErrInternalServerError.Error())
	}

	return &api.GetUserResponse{
		Data: &api.User{
			Id:      user.Data.ID,
			Name:    user.Data.Name,
			Surname: user.Data.Surname,
		},
	}, nil
}
