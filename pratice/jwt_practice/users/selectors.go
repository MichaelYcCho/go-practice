package users

import "gorm.io/gorm"

type Selector interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type selector struct {
	db *gorm.DB
}

func NewSelector(db *gorm.DB) *selector {
	return &selector{db}
}

func (s *selector) Save(user User) (User, error) {
	err := s.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *selector) FindByEmail(email string) (User, error) {
	var user User
	err := s.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *selector) FindByID(ID int) (User, error) {
	var user User
	err := s.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *selector) Update(user User) (User, error) {
	err := s.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
