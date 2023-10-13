package user_dtos

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
