package main

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/auth"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/database"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/handler"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/middleware"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitialDB()

	// Todo: 마이그레이션 폴더로 이동시키는것이 나을까?
	modelList := []interface{}{
		&users.User{},
	}
	db.AutoMigrate(modelList...)

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

}
