package types

import (
	"database/sql"
	"encoding/json"
)

type NullInt32 struct {
	sql.NullInt32
}

func (v NullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullInt32) UnmarshalJSON(data []byte) error {
	var x *int32
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func NewNullInt32(n *int32) NullInt32 {
	return NullInt32{
		sql.NullInt32{
			Int32: *n,
			Valid: n != nil,
		},
	}
}
