package member

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user already exists")
	ErrUserNameIsRequired       = errors.New("user name is required")
	ErrMembershipTypeIsRequired = errors.New("membership type is required")
	ErrInvalidMembershipType    = errors.New("choose membership type : naver, payco, toss")
	ErrUserIDNotFound           = errors.New("user id not found")
	ErrUserIDIsRequired         = errors.New("user id is required")
)
