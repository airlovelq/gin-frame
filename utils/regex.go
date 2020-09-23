package utils

import (
	"regexp"
)

func IsEmail(str string) bool {
	is, err := regexp.MatchString(`^[\.A-Za-z0-9]+@[\.a-zA-Z0-9_-]+$`, str)
	return err == nil && is
}

func IsPhone(str string) bool {
	is, err := regexp.MatchString(`^(1[3-9])\d{9}$`, str)
	return err == nil && is
}
