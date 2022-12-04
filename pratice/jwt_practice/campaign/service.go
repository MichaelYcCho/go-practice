package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
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
