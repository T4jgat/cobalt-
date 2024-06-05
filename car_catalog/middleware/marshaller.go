package middleware

import "time"

type Datetime time.Time

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

func (d Datetime) MarshalJSON() ([]byte, error) {
	// TODO implement
	t := time.Time(d)
	formatted := t.Format("2 Jan 2006 3:04PM")

	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}
