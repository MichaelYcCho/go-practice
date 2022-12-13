package transaction

import "os/user"

type InputGetCampaignTransaction struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
