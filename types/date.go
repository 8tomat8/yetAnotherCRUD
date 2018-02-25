package types

import (
	"time"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(data []byte) error {
	date, err := time.Parse("\"02-01-2006\"", string(data))
	if err != nil {
		return err
	}
	*t = Date{date}
	return nil
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(t.Format("\"02-01-2006\"")), nil
}
