package controllers

import (
	"crud-go-boilerplate-fiber/app/models/requests"
	"crud-go-boilerplate-fiber/app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserControllers interface {
	Login(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
type userControllers struct {
	usrService services.UserService
}

func NewUserControllers(s services.UserService) UserControllers {
	return &userControllers{usrService: s}
}

func (c *userControllers) Login(ctx *fiber.Ctx) error {
	var params requests.UserLoginRequest
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	res, err := c.usrService.Login(params)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})

	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}
func (c *userControllers) Create(ctx *fiber.Ctx) error {
	var params requests.UserRequest
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	res, err := c.usrService.Create(params)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})

	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}
