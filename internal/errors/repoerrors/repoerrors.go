package repoerrors

import "errors"

var ErrRepoUniqueEmailViolation = errors.New("unique email violation")
var ErrRepoUniqueUserNameViolation = errors.New("unique username violation")
var ErrRepoRecordNotFound = errors.New("user with that username does not exist") // to check
