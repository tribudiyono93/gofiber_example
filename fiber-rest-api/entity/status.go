package entity

const (
	UserNotFound           = 1000
	UserAlreadyExist       = 1001
	EmailAlreadyUsed       = 1002
	InvalidEmailOrPassword = 1003
)

var StatusText = map[int]string{
	UserNotFound:           "user not found",
	UserAlreadyExist:       "user already exist",
	EmailAlreadyUsed:       "email already used",
	InvalidEmailOrPassword: "invalid email or password",
}
