package usecase

import (
	"simple-attendance/internal/user/models/domain"
	"simple-attendance/internal/user/models/entity"
)

type UserUseCase interface {
	AddUser(req domain.ReqAddUser) (domain.UserId, error)
	GetUserById(id string) (domain.ResponseUser, error)
	GetUsersByUsername(username string) ([]domain.ResponseUser, error)

	VerifyUserCredential(req entity.User) (entity.User, error)
}
