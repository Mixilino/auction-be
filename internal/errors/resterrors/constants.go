package resterrors

import "errors"

var ConstPasswordNotStrongEnough = errors.New("password must contain at least 8 characters")
