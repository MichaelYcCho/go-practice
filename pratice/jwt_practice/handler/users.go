package handler

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/helper"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", 400, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

}
