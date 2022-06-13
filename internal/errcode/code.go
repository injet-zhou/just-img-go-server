package errcode

const (
	ErrWrongPassword = iota + 1001
	ErrUserNameOrEmailRequired
	ErrWrongUsername
	ErrLoginFailTooManyTimes
	ErrUserNotExist
	DBErr
	ErrTokenUnauthorized
	ErrTokenExpired
)
