package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/akash/go-user-api/config"
	"github.com/akash/go-user-api/internal/models"
	"github.com/akash/go-user-api/internal/repository"
	"github.com/akash/go-user-api/internal/service"
)

func CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}
	if err := service.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.CreateUser(
		config.DB,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	user, err := repository.GetUserByID(
		config.DB,
		id,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	dob, err := time.Parse(
		time.RFC3339,
		user.DOB,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB,
		Age:  service.CalculateAge(dob),
	}

	return c.JSON(response)
}
func GetAllUsers(c *fiber.Ctx) error {

	users, err := repository.GetAllUsers(config.DB)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []models.UserResponse

	for _, user := range users {

		dob, err := time.Parse(
			time.RFC3339,
			user.DOB,
		)

		if err != nil {
			continue
		}

		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.DOB,
			Age:  service.CalculateAge(dob),
		})
	}

	return c.JSON(response)
}
func UpdateUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	if err := service.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.UpdateUser(
		config.DB,
		id,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	err = repository.DeleteUser(
		config.DB,
		id,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.SendStatus(204)
}
