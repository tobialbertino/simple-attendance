package domain

type ResponseUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type UserData struct {
	User ResponseUser `json:"user"`
}

type UsersData struct {
	User []ResponseUser `json:"users"`
}
