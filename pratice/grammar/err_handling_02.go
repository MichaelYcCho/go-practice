package grammar

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("Login failed: %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("File not found: %s", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) error {
	if uid != "admin" || pwd != "admin" {
		return errors.New("invalid login")
	}
	return nil
}

func getData(file string) ([]byte, error) {
	return nil, nil
}
