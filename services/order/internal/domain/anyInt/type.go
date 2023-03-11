package anyInt

import "strconv"

type Int int64

func (i Int) String() string {
	return strconv.FormatUint(uint64(i), 10)
}
func New(val int64) *Int {
	i := Int(val)
	return &i
}
func NewStr(str string) (Int, error) {
	if val, err := strconv.Atoi(str); err != nil {
		return 0, err
	} else {
		i := Int(int64(val))
		return i, nil
	}
}
func (i Int) Int() int {
	return int(i)
}
