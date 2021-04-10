package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/connection"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/request"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/response"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/util"
	"log"
	"net/http"
	"os"
	"time"
)

func Register(c *fiber.Ctx) error {
	req := new(request.Register)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(response.StatusText[response.InternalServerError])
	}

	var user entity.User
	connection.DB.Where("email = ?", req.Email).First(&user)
	if user.ID != "" {
		return c.Status(http.StatusBadRequest).JSON(response.Error{
			Code: response.UserAlreadyExist, Message: response.StatusText[response.UserAlreadyExist]})
	}

	base := entity.Base{
		CreatedBy: req.Email, CreatedAt: time.Now(), UpdatedBy: req.Email, UpdatedAt: time.Now(),
	}
	hash, err := util.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(response.StatusText[response.InternalServerError])
	}
	user = entity.User{
		ID: utils.UUIDv4(), Email: req.Email, Password: hash, Name: req.Name, Base : base,
	}
	connection.DB.Create(&user)

	var userModuleRoles []entity.UserModuleRole
	for _, v := range req.UserModuleRoles {
		userModuleRole := entity.UserModuleRole{
			Email: req.Email, Module: v.Module, Role: v.Role, Base: base,
		}
		connection.DB.Create(userModuleRole)
		userModuleRoles = append(userModuleRoles, userModuleRole)
	}

	return c.Status(http.StatusOK).JSON(response.UserDetail{User: user, UserModuleRoles: userModuleRoles})
}

func Login(c *fiber.Ctx) error {
	req := new(request.Login)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(response.StatusText[response.InternalServerError])
	}

	var user entity.User
	connection.DB.Where("email = ?", req.Email).First(&user)
	if user.ID == "" {
		return c.Status(http.StatusBadRequest).JSON(response.Error{
			Code: response.UserNotFound, Message: response.StatusText[response.UserNotFound]})
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(http.StatusBadRequest).JSON(response.Error{
			Code: response.InvalidEmailOrPassword, Message: response.StatusText[response.InvalidEmailOrPassword]})
	}

	var userModuleRoles []entity.UserModuleRole
	connection.DB.Where("email = ?", user.Email).Find(&userModuleRoles)

	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(60 * time.Minute)
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := j.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(response.StatusText[response.InternalServerError])
	}

	return c.Status(http.StatusOK).JSON(response.UserDetail{User: user, UserModuleRoles: userModuleRoles, Token: token})
}
