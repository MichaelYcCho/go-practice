package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	GenerateToken(userID int) (string, error)
}

// 메서드를 만들기위해 빈 타입을 선언했다고 이해하는데 더 나은 구조는 없을까?
type jwtService struct {
}

var SECRET_KEY = []byte("SECRETKEY")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
