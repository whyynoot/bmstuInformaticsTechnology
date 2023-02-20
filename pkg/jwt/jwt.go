package jwt

import (
	"bmstuInformaticsTechnologies/internal/user_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var SecretKey = []byte("n5VFbBxdSjGiymiKaZvGeKzzOlMRR8G6")

type Helper struct {
	logger logging.LoggerInterface
}

func (h Helper) GenerateToken(user *user_service.User) (*Authorization, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	return &Authorization{
		Token:        t,
		RefreshToken: rt,
	}, nil
}

func (h Helper) RefreshToken(authorization *Authorization) (Authorization, error) {
	//TODO implement me
	panic("implement me")
}

func NewHelper(loggerInterface logging.LoggerInterface) *Helper {
	return &Helper{
		logger: loggerInterface,
	}
}
