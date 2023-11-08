package controllers

import (
	"crud-go-boilerplate-fiber/app/models/requests"
	"crud-go-boilerplate-fiber/app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	Save(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
type productController struct {
	ps services.ProductService
}

func NewProductController(ps services.ProductService) ProductController {
	return &productController{ps: ps}
}
func (c *productController) Save(ctx *fiber.Ctx) error {
	var params requests.InputProduct
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	res, err := c.ps.Create(params)
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
func (c *productController) Update(ctx *fiber.Ctx) error {
	var params requests.UpdateProduct
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	uid := ctx.Query("uid")
	params.ProductUID = uid

	res, err := c.ps.Update(params)
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
