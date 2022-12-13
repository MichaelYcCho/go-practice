package transaction

import "gorm.io/gorm"

type Selector interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
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
