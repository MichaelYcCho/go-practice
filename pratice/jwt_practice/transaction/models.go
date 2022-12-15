package transaction

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/campaign"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       users.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
