package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/connection"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/util"
	"log"
	"net/http"
)

func HasModuleAndRole(moduleRoles map[string][]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email, err := util.GetEmailFromToken(c.Get(authorization))
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString(http.StatusText(http.StatusInternalServerError))
		}

		var userModuleRoles []entity.UserModuleRole
		connection.DB.Where("email = ?", email).Find(&userModuleRoles)

		allow := false

		for _, v := range userModuleRoles {
			if val, ok := moduleRoles[v.Module]; ok {
				if exist := find(val, v.Role); exist {
					allow = true
					break
				}
			}
		}

		if !allow {
			return c.Status(http.StatusForbidden).SendString(http.StatusText(http.StatusForbidden))
		}

		return c.Next()
	}
}

func find(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}
