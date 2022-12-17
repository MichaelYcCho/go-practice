package company

import "gorm.io/gorm"

type Selector interface {
	CreateCompany(company Company) (Company, error)

}

type selector struct {
	db *gorm.DB
}

func NewSelector(db *gorm.DB) *selector {
	return &selector{db}
}

func (s *selector) CreateCompany(company Company) (Company, error) {
	err := s.db.Create(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}