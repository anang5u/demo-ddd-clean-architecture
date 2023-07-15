package controller

import (
	"demo-ddd-clean-architecture/app/model"

	"github.com/gofiber/fiber/v2"
)

// SendSuccess
func SendSuccess(c *fiber.Ctx, data interface{}) error {
	return c.JSON(model.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}

// SendError
func SendError(c *fiber.Ctx, err error) error {
	return c.JSON(model.ApiResponse{
		Code:   201,
		Status: "Error",
		Data:   err.Error(),
	})
}
