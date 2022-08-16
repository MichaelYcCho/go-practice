package grammar

import "fmt"

type StatusCustomErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusCustomErr) Error() string {
	return se.Message
}

func (se StatusCustomErr) Unwrap() error {
	return se.Err
}

func LoginAndGetData02(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusCustomErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("Login failed: %s", uid),
			Err:     err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusCustomErr{
			Status:  NotFound,
			Message: fmt.Sprintf("File not found: %s", file),
			Err:     err,
		}
	}
	return data, nil

}
