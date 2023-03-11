package phone

import (
	"fmt"
	"regexp"
)

var (
	pattern       = regexp.MustCompile("^[+]?[(]?([0-9]{3})?[)]?[-\\s.]?[0-9]{3}[-\\s.]?[0-9]{4,6}$")
	ErrWrongPhone = fmt.Errorf("phone number is wrong")
)

type Phone struct {
	PhoneNumber string
}

func (p Phone) String() string {
	return p.PhoneNumber
}
func New(phone string) (Phone, error) {
	if !pattern.MatchString(phone) {
		return Phone{}, ErrWrongPhone
	}
	return Phone{
		PhoneNumber: phone,
	}, nil
}
