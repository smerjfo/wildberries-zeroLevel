package anyString

import (
	"fmt"
)

const (
	MaxLength = 200
)

var ErrWrongLength = fmt.Errorf("string length must be less then %d", MaxLength)

type AnyString string

func (s AnyString) String() string {
	return string(s)
}
func New(anySting string) (AnyString, error) {
	if len([]rune(anySting)) > MaxLength {
		return "", ErrWrongLength
	}
	s := AnyString(anySting)
	return s, nil
}
