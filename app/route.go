package server

import (
	"crud-go-boilerplate-fiber/app/controllers"
	"crud-go-boilerplate-fiber/app/libraries/middleware"
	"crud-go-boilerplate-fiber/app/repository"
	"crud-go-boilerplate-fiber/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppRoute(db *gorm.DB, app *fiber.App) {
	baseRepo := repository.NewBaseRepository(db)
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(baseRepo, productRepo)
	productController := controllers.NewProductController(productService)
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(baseRepo, userRepo)
	userController := controllers.NewUserControllers(userService)
	auth := middleware.NewAuth()

	user := app.Group("/api/v1")
	user.Post("/login", userController.Login)
	user.Post("/user", userController.Create)

	product := app.Group("/api/v1", auth.CheckAuth)

	product.Post("/product", productController.Save)
	product.Post("/product/:uid", productController.Update)

}
