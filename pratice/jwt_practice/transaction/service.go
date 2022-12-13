package transaction

import (
	"errors"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input InputGetCampaignTransaction) ([]Transaction, error)
}

type service struct {
	selector         Selector
	campaignSelector campaign.Selector
}

func NewService(selector Selector, campaignSelector campaign.Selector) *service {
	return &service{selector, campaignSelector}
}

func (s *service) GetTransactionsByCampaignID(input InputGetCampaignTransaction) ([]Transaction, error) {
	campaign, err := s.campaignSelector.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.selector.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
