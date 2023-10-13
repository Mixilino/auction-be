package user_dtos

type SignUpRequest struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
