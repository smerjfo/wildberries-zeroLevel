package email

import (
	"fmt"
	"regexp"
)

const Length = 100

var (
	pattern    = regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+")
	ErrPattern = fmt.Errorf("email is not valid")
	ErrLength  = fmt.Errorf("email length must be less then %d characters", Length)
)

type Email struct {
	Email string
}

func (e Email) String() string {
	return e.Email
}
func New(email string) (Email, error) {
	if len(email) > Length {
		return Email{}, ErrLength
	}
	if !pattern.MatchString(email) {
		return Email{}, ErrPattern
	}
	return Email{Email: email}, nil
}
