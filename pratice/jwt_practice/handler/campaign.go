package handler

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/campaign"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {

	// strconv -> 문자열을 받아 int로 파싱해줌
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)

}
