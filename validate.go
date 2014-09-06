package validate

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

func StringExist(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func Int64Exist(input int64) (err error) {
	if input == 0 {
		err = NewError("must Exist")
	}
}

func AlphaNumeric(input string) (err error) {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(input) {
		err = NewError("must only contain letters, numbers, or an underscore")
	}
	return
}

func Username(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func Password(input string) (err error) {
	input = strings.TrimSpace(input)
	[]rune(input)
	return
}

func Email(input string) (err error) {
	input = strings.TrimSpace(input)
	_, err = mail.ParseAddress(input)
	return
}

func Url(input string) (err error) {
	input = strings.TrimSpace(input)
	_, err = url.Parse(input)
	return
}

// Valid javascript date number
func JsDate(input int64) (err error) {

	return
}

// func PhoneNumber(input string) (err error) {
// 	input = strings.TrimSpace(input)
// 	return
// }

// func ZipCode(input string) (err error) {
// 	input = strings.TrimSpace(input)
// 	return
// }

// func Country(input string) (err error) {
// 	input = strings.TrimSpace(input)
// 	return
// }

// func City(input string) (err error) {
// 	input = strings.TrimSpace(input)
// 	return
// }

// func StateOrProvince(input string) (err error) {
// 	input = strings.TrimSpace(input)
// 	return
// }
