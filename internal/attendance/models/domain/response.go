package domain

type ResAttendance struct {
	Id       string  `json:"id"`
	UserId   string  `json:"user_id"`
	Activity *string `json:"activity"`
	Location *string `json:"location"`
	CheckIn  *int64  `json:"check_in"`
	CheckOut *int64  `json:"check_out"`
}

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IdAttendance struct {
	ID string `json:"id_attendance"`
}
