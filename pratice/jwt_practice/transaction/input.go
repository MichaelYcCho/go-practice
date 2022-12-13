package transaction

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
)

type InputGetCampaignTransaction struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}
