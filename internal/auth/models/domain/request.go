package domain

type ReqLoginUser struct {
	Username  string `json:"username" validate:"required"`
	Passwword string `json:"password" validate:"required,min=6"`
}

type ReqRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
