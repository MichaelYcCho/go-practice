package routers

import (
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/database"
	"github.com/MichaelYcCho/go-practice/pratice/simple-book/handelrs"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handelrs.APIEnv{
		DB: database.GetDB(),
	}

	r.GET("", api.GetBooks)
	// r.GET("/:id", api.GetBook)
	// r.POST("", api.CreateBook)
	// r.PUT("/:id", api.UpdateBook)
	// r.DELETE("/:id", api.DeleteBook)

	return r
}
