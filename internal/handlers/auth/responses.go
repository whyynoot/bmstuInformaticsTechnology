package auth

import (
	"bmstuInformaticsTechnologies/internal/user_service"
	"bmstuInformaticsTechnologies/pkg/jwt"
)

type AuthorizationResponse struct {
	*jwt.Authorization
}

// UserLoginRequest is a struct provided for person to request auth, captcha can be added
type UserLoginRequest struct {
	user_service.UserDTO
}
