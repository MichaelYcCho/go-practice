package campaign

import "gorm.io/gorm"

type Selector interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type selector struct {
	db *gorm.DB
}

func NewSelector(db *gorm.DB) *selector {
	return &selector{db}
}

func (s *selector) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	// Django의 Prefetch와 같은개념 2번의 쿼리로 선행 쿼리 결과값에 매칭
	// Campaign 레코드마다 CampaignImages가 있다면 campaignImages에 데이터가 담김
	// 쿼리문 확인을 위해 Debug() 사용
	err := s.db.Debug().Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *selector) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := s.db.Debug().Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
