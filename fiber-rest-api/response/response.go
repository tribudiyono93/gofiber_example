package response

import "github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"

const (
	UserNotFound           = 1000
	UserAlreadyExist       = 1001
	EmailAlreadyUsed       = 1002
	InvalidEmailOrPassword = 1003
	InternalServerError    = 9999
)

var StatusText = map[int]string{
	UserNotFound:           "user not found",
	UserAlreadyExist:       "user already exist",
	EmailAlreadyUsed:       "email already used",
	InvalidEmailOrPassword: "invalid email or password",
	InternalServerError:    "internal server error",
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pageable struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}

type UserDetail struct {
	User            entity.User             `json:"user"`
	UserModuleRoles []entity.UserModuleRole `json:"userModuleRoles"`
	Token           string                  `json:"token,omitempty"`
}
