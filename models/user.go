package models

import (
	"auction-be/constants/regexconstants"
	"auction-be/utils/regexutils"
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Password  []byte    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) Validate() error {
	if user.UserName == "" {
		return errors.New("username must not be empty")
	}

	if err := regexutils.ValidateStringWithRegex(user.UserName, regexconstants.ValidUsername); err != nil {
		return errors.New("username must contain 6-20 alphanumeric characters, must start with character" + err.Error())
	}

	if len(user.Password) == 0 {
		return errors.New("password must not be empty")
	}

	if user.FirstName == "" {
		return errors.New("firstName must not be empty")
	}

	if user.LastName == "" {
		return errors.New("lastName must not be empty")
	}

	if err := regexutils.ValidateStringWithRegex(user.Email, regexconstants.ValidEmail); err != nil {
		return errors.New("email is not valid" + err.Error() + user.Email)
	}
	return nil
}
