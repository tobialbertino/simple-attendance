package domain

type ReqAttendance struct {
	Id       string  `json:"id" validate:"required"`
	UserId   string  `json:"user_id" validate:"required"`
	Activity *string `json:"activity" validate:"required"`
	Location *string `json:"locatioon" validate:"required"`
	CheckIn  *int64  `json:"check_in" validate:"required"`
	CheckOut *int64  `json:"check_out" validate:"required"`
}

type ReqAddAttendance struct {
	Id     string `json:"id"`
	UserId string `json:"user_id" validate:"required"`
}

type ReqUpdate struct {
	Id       string  `json:"id" validate:"required"`
	Activity *string `json:"activity" validate:"required"`
	Location *string `json:"location" validate:"required"`
}

type ReqCheckIn struct {
	Id      string `json:"id" validate:"required"`
	CheckIn *int64 `json:"check_in"`
}

type ReqCheckOut struct {
	Id       string `json:"id" validate:"required"`
	CheckOut *int64 `json:"check_out"`
}

type ReqId struct {
	Id string `json:"id" validate:"required"`
}

type ReqGetAllByUserId struct {
	UserId string `json:"user_id" validate:"required"`
}
