package jwt

import "bmstuInformaticsTechnologies/internal/user_service"

type HelperInterface interface {
	GenerateToken(user *user_service.User) (*Authorization, error)
	RefreshToken(authorization *Authorization) (Authorization, error)
}
