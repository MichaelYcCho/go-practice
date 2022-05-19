package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {

	})

	//post user sign in
	r.POST("/user/signin", func(c *gin.Context) {

	})

	return r.Run(address)
}
