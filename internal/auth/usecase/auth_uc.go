package usecase

import "simple-attendance/internal/auth/models/domain"

type AuthUseCase interface {
	AddRefreshToken(req domain.ReqLoginUser) (domain.ResToken, error)
	VerifyRefreshToken(req domain.ReqRefreshToken) (domain.ResToken, error)
	DeleteRefreshToken(req domain.ReqRefreshToken) (domain.ResToken, error)
}
