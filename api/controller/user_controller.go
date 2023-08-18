package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/errs"
	"github.com/karlosdaniel451/go-rest-api-template/usecase"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return UserController{userUseCase: useCase}
}

// Create a new User.
// @Description Create a new User.
// @Summary Create a new User.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Router /users [post]
func (controller UserController) Create(c *fiber.Ctx) error {
	var newUser model.User

	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"detail": "invalid user data: " + err.Error(),
		})
	}

	newUserAllData, err := controller.userUseCase.Create(&newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUserAllData)
}

// Delete a User.
// @Description Delete a User by its id.
// @Summary Delete a User and, in case there is no User with the given ID,
// returns a 404 not found error.
// @Tags Users
// @Produce json
// @Param id path int true "Id of the User be deleted"
// @Success 204
// @Failure 404
// @Router /users/{id} [delete]
func (controller UserController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	err = controller.userUseCase.DeleteById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Get a User by its id.
// @Description Get a User by its id.
// @Summary Get a User by its id.
// @Tags Users
// @Produce json
// @Success 200 {object} model.User
// @Failure 404
// @Router /users/{id} [get]
func (controller UserController) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer greater than 0",
		})
	}

	task, err := controller.userUseCase.GetById(uint(id))
	if err != nil {
		if errors.As(err, &errs.NotFoundError{}) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.JSON(task)
}

// Get all Users.
// @Description Get all Users.
// @Summary Get all Users.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func (controller UserController) GetAll(c *fiber.Ctx) error {
	tasks, err := controller.userUseCase.GetAll()
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.JSON(tasks)
}
