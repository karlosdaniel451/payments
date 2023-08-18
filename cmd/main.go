package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
	"github.com/karlosdaniel451/go-rest-api-template/api/middleware"
	"github.com/karlosdaniel451/go-rest-api-template/api/router"
	"github.com/karlosdaniel451/go-rest-api-template/db"
	_ "github.com/karlosdaniel451/go-rest-api-template/docs"
	"github.com/karlosdaniel451/go-rest-api-template/repository"
	"github.com/karlosdaniel451/go-rest-api-template/usecase"
)

var port = os.Getenv("API_PORT")

// @title Payments
// @version 0.0.1
// @description Payments RESTful Web Service with Go, Fiber and GORM
func main() {
	var _ usecase.UserUseCase = usecase.UserUseCaseImpl{}
	var _ repository.UserRepository = repository.UserRepositoryDB{}

	var _ usecase.TransactionUseCase = usecase.TransactionUseCaseImpl{}
	var _ repository.TransactionRepository = repository.TransactionRepositoryDB{}

	app := fiber.New(fiber.Config{
		AppName:           "Payments RESTful Web Service with Go, Fiber and GORM",
		EnablePrintRoutes: true,
	})

	middleware.Setup(app)

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepositoryDB(db.DB)
	userUseCase := usecase.NewUserUseCaseImpl(userRepository)
	userController := controller.NewUserController(userUseCase)

	router.Setup(app, &userController)

	log.Fatal(app.Listen(":" + port))
}
