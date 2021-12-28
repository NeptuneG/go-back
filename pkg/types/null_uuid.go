package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

type NullUUID struct {
	uuid.NullUUID
}

func (s NullUUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.UUID)
}

func (s *NullUUID) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	s.UUID, _ = uuid.Parse(str)
	s.Valid = str != ""
	return nil
}

func NewNullUUID(id *uuid.UUID) NullUUID {
	if id == nil {
		return NullUUID{
			uuid.NullUUID{
				Valid: false,
			},
		}
	} else {
		return NullUUID{
			uuid.NullUUID{
				Valid: true,
				UUID:  *id,
			},
		}
	}
}
