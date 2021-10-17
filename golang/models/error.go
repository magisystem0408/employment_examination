package models

type InvaildError struct {
	Massage string
	ErrCode int
}

func (e InvaildError) Error() string {
	return e.Massage
}
