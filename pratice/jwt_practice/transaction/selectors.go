package transaction

import "gorm.io/gorm"

type Selector interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
}

type selector struct {
	db *gorm.DB
}

func NewSelector(db *gorm.DB) *selector {
	return &selector{db}
}

func (s *selector) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := s.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *selector) GetByUserID(userID int) ([]Transaction, error) {
	var transaction []Transaction
	err := s.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
