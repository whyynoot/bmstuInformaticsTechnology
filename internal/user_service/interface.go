package user_service

type UserServiceInterface interface {
	GetUserByUsernameAndPassword(username, password string) (*User, error)
	Create(dto UserDTO) (*User, error)
}
