package users

import (
	"errors"
	"fmt"
	"github.com/MichaelYcCho/go-practice/pratice/jwt_practice/utils"
	"golang.org/x/crypto/bcrypt"
	"path/filepath"
	"strings"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	MakePathStr(name string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	_, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}
	return user, nil
}

func (s *service) fileNameUsed(name string) (bool, error) {
	matches, err := filepath.Glob("images/*" + name)
	if err != nil {
		return false, err
	}
	return len(matches) > 0, nil
}

func (s *service) MakePathStr(imageName string) (string, error) {
	path := fmt.Sprintf("images/%s", imageName)

	for true {
		match, err := s.fileNameUsed(imageName)
		if err != nil {
			return "", err
		}
		if match {
			randomStr := utils.RandString(2)
			fileSlice := strings.Split(imageName, ".")
			fileExtension := fileSlice[len(fileSlice)-1]
			fileName := fmt.Sprintf(strings.Join(fileSlice[:len(fileSlice)-1], ".")+"_%s", randomStr)
			imageName = fmt.Sprintf("%s.%s", fileName, fileExtension)
		} else {
			break
		}

	}
	path = fmt.Sprintf("images/%s", imageName)
	return path, nil
}
