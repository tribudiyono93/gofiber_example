package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/connection"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/response"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/util"
	"log"
	"net/http"
)

func Profile(c *fiber.Ctx) error {
	email, err := util.GetEmailFromToken(c.Get("Authorization"))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString(http.StatusText(http.StatusInternalServerError))
	}

	var user entity.User
	connection.DB.Where("email = ?", email).First(&user)
	if user.ID == "" {
		return c.Status(http.StatusBadRequest).JSON(response.Error{
			Code: response.UserNotFound, Message: response.StatusText[response.UserNotFound]})
	}

	var userModuleRoles []entity.UserModuleRole
	connection.DB.Where("email = ?", user.Email).Find(&userModuleRoles)

	return c.Status(http.StatusOK).JSON(response.UserDetail{User: user, UserModuleRoles: userModuleRoles})
}
