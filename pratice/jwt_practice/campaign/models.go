package campaign

import (
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/users"
	"time"
)

// CampaignImages와 Campaign을 1:N 관계로 설정
type Campaign struct {
	ID                int
	UserID            int
	Name              string
	ShortDescriptions string
	Descriptions      string
	Perks             string
	BackerCount       int
	GoalAmount        int
	CurrentAmount     int
	Slug              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CampaignImages    []CampaignImage
	User              users.User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
