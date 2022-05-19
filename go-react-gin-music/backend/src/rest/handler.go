package rest

import (
	"github.com/MichaelYcCho/go-practice/go-react-gin-music/backend/src/dblayer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db dblayer.DBLayer
}

func Newhandler() (*Handler, error) {
	return new(Handler), nil
}

type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}
