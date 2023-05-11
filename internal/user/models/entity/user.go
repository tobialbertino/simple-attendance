package entity

import "simple-attendance/internal/user/models/domain"

type ListUser []User
type User struct {
	Id        string
	Username  string
	Passwword string
	FullName  string
}

func (dt *User) ToDomain() domain.ResponseUser {
	return domain.ResponseUser{
		Id:       dt.Id,
		Username: dt.Username,
		FullName: dt.FullName,
	}
}

func (ldt *ListUser) ToDomain() []domain.ResponseUser {
	var result []domain.ResponseUser = make([]domain.ResponseUser, 0)
	for _, dt := range *ldt {
		result = append(result, dt.ToDomain())
	}

	return result
}
