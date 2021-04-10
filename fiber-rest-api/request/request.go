package request

type UserModuleRoleRequest struct {
	Module string `json:"module" binding:"required"`
	Role   string `json:"role" binding:"required"`
}

type Register struct {
	Email           string                  `json:"email" binding:"required"`
	Password        string                  `json:"password" binding:"required"`
	Name            string                  `json:"name" binding:"required"`
	UserModuleRoles []UserModuleRoleRequest `json:"userModuleRoles"`
}

type Login struct {
}

type CreateBook struct {
}
