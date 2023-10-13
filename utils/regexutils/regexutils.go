package regexutils

import (
	"errors"
	"regexp"
)

func ValidateStringWithRegex(input string, regexPattern string) error {
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return err
	}

	if !regex.MatchString(input) {
		return errors.New("string does not match the regex pattern")
	}

	return nil
}
