package user_service

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
