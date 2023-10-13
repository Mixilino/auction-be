package conversions

import (
	"auction-be/dtos/user_dtos"
	"auction-be/models"
)

func UserModelToUserResponse(user models.User) *user_dtos.Response {
	return &user_dtos.Response{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreateAt:  user.CreateAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserSignUpRequestToUserModel(userDto user_dtos.SignUpRequest, password []byte) *models.User {
	return &models.User{
		UserName:  userDto.UserName,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		Password:  password,
	}
}
