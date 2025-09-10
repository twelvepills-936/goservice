package models

import (
	ucModels "gitlab16.skiftrade.kz/templates/go/internal/usecase/models"
	api "gitlab16.skiftrade.kz/templates/go/pkg/api"
)

func ToProtoUser(u ucModels.User) *api.User {
	userProto := &api.User{}
	userProto.SetId(u.ID)
	userProto.SetName(u.Name)
	userProto.SetSurname(u.Surname)
	return userProto
}
