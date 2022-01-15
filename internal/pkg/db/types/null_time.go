package types

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	var tm time.Time
	if err := json.Unmarshal(data, &tm); err != nil {
		return err
	}
	t.Time = tm
	t.Valid = !tm.IsZero()
	return nil
}

func NewNullTime(t time.Time) NullTime {
	return NullTime{
		sql.NullTime{
			Time:  t,
			Valid: !t.IsZero(),
		},
	}
}
