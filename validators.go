package validators

import (
	"errors"
	"strings"
)

var (
	whitespace = " \n"
)

func Email(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func PhoneNumber(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

// Valid javascript date number
func JsDate(input int64) (err error) {

	return
}

func Username(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

// Password must
func Password(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func StringExist(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func Url(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func ZipCode(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func Country(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func City(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}

func StateOrProvince(input string) (err error) {
	input = strings.TrimSpace(input)
	return
}
