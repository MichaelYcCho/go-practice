package handler

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/company"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type companyHandler struct {
	service company.Service
}

func NewCompanyHandler(service company.Service) *companyHandler {
	return &companyHandler{service}
}

func (h *companyHandler) CreateCompany(c *gin.Context) {
	var input company.InputCompany

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCompany, err := h.service.CreateCompany(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIResponse("Success to create company", http.StatusOK, "success", company.FormatCompany(newCompany))
	c.JSON(http.StatusOK, response)



}
