package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/auth"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/campaign"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/company"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/database"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/handler"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/middleware"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/transaction"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitialDB()

	// Todo: 마이그레이션 폴더로 이동시키는것이 나을까?
	modelList := []interface{}{
		&users.User{},
		&campaign.Campaign{},
		&campaign.CampaignImage{},
		&transaction.Transaction{},
		&company.Company{},
	}
	db.AutoMigrate(modelList...)
	authService := auth.NewService()

	userRepository := users.NewSelector(db)
	userService := users.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewSelector(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	companyRepository := company.NewSelector(db)
	companyService := company.NewService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)


	transactionRepository := transaction.NewSelector(db)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)

	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransactions)

	api.POST("/companies", companyHandler.CreateCompany)

	router.Run()

}
