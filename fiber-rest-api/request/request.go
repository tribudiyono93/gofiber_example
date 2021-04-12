package request

import "github.com/go-playground/validator"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func Validate(reqBody interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(reqBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type UserModuleRoleRequest struct {
	Module string `json:"module" validate:"required"`
	Role   string `json:"role" validate:"required"`
}

type Register struct {
	Email           string                  `json:"email" validate:"required"`
	Password        string                  `json:"password" validate:"required"`
	Name            string                  `json:"name" validate:"required"`
	UserModuleRoles []UserModuleRoleRequest `json:"userModuleRoles"`
}

type Login struct {
	Email           string                  `json:"email" validate:"required"`
	Password        string                  `json:"password" validate:"required"`
}

type RefreshJWTToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type CreateBook struct {
}
