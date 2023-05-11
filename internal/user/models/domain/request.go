package domain

type ReqAddUser struct {
	Username  string `json:"username" validate:"required"`
	Passwword string `json:"password" validate:"required,min=6"`
	FullName  string `json:"fullname" validate:"required"`
}

type UserId struct {
	UserId string `json:"userId"`
}
