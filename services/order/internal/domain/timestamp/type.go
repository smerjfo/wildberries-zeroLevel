package timestamp

import "time"

type Timestamp struct {
	Timestamp time.Time
}

func (t Timestamp) String() string {
	return t.Timestamp.String()
}
func New(timestamp string) (Timestamp, error) {
	val, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timestamp)
	if err != nil {
		return Timestamp{}, err
	}
	return Timestamp{
		val,
	}, nil
}
func (t Timestamp) Time() time.Time {
	return t.Timestamp
}
