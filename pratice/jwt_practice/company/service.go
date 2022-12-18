package company

type Service interface {
	CreateCompany(input InputCompany) (Company, error)

}

type service struct {
	selector Selector
}

func NewService(selector Selector) *service {
	return &service{selector}
}

func (s *service) CreateCompany(input InputCompany) (Company, error) {
	company := Company{}
	company.Name = input.Name
	company.CEO = input.CEO
	company.Revenue = input.Revenue

	newCompany, err := s.selector.CreateCompany(company)
	if err != nil {
		return newCompany, err
	}
	return newCompany, nil
}