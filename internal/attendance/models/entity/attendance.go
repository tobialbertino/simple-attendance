package entity

import "simple-attendance/internal/attendance/models/domain"

type ListAttendances []Attendance
type Attendance struct {
	Id       string
	UserId   string
	Activity *string
	Location *string
	CheckIn  *int64
	CheckOut *int64
}

func (dt *Attendance) ToDomain() domain.ResAttendance {
	return domain.ResAttendance{
		Id:       dt.Id,
		UserId:   dt.UserId,
		Activity: dt.Activity,
		Location: dt.Location,
		CheckIn:  dt.CheckIn,
		CheckOut: dt.CheckOut,
	}
}

func (ldt *ListAttendances) ToDomain() []domain.ResAttendance {
	var listResult []domain.ResAttendance = make([]domain.ResAttendance, 0)
	for _, dt := range *ldt {
		listResult = append(listResult, dt.ToDomain())
	}
	return listResult
}
