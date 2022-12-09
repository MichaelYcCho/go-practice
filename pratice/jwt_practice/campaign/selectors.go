package campaign

import "gorm.io/gorm"

type Selector interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Create(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
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

func (s *selector) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := s.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *selector) Create(campaign Campaign) (Campaign, error) {
	err := s.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *selector) Update(campaign Campaign) (Campaign, error) {
	err := s.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *selector) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := s.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (s *selector) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	/* Update CampaignImage set is_primary = false where campaign_id = 1 */
	err := s.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
