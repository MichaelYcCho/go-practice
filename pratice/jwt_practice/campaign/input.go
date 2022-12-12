package campaign

import "github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"

type InputCampaignDetail struct {
	ID int `uri:"id" binding:"required"`
}

type InputCampaign struct {
	Name              string     `json:"name" binding:"required"`
	ShortDescriptions string     `json:"short_descriptions" binding:"required"`
	Descriptions      string     `json:"descriptions" binding:"required"`
	GoalAmount        int        `json:"goal_amount" binding:"required"`
	Perks             string     `json:"perks" binding:"required"`
	User              users.User //
}

type InputCreateCampaignImage struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       users.User
}
