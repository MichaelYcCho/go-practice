package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input InputCampaignDetail) (Campaign, error)
	CreateCampaign(input InputCampaign) (Campaign, error)
	UpdateCampaign(inputID InputCampaignDetail, inputData InputCampaign) (Campaign, error)
	SaveCampaignImage(input InputCreateCampaignImage, fileLocation string) (CampaignImage, error)
}

type service struct {
	selector Selector
}

func NewService(selector Selector) *service {
	return &service{selector}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.selector.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.selector.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input InputCampaignDetail) (Campaign, error) {
	campaign, err := s.selector.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input InputCampaign) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescriptions = input.ShortDescriptions
	campaign.Descriptions = input.Descriptions
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.User.ID = input.User.ID

	// Slug 라이브러리 (gosimple)
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate) // Name Campaign + ID

	newCampaign, err := s.selector.Create(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID InputCampaignDetail, inputData InputCampaign) (Campaign, error) {
	campaign, err := s.selector.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	campaign.Name = inputData.Name
	campaign.ShortDescriptions = inputData.ShortDescriptions
	campaign.Descriptions = inputData.Descriptions
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.selector.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input InputCreateCampaignImage, fileLocation string) (CampaignImage, error) {

	campaign, err := s.selector.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.selector.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.selector.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil

}
