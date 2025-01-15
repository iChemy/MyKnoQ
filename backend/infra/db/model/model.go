package model

import (
	"database/sql/driver"

	"github.com/gofrs/uuid"
)

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

func (nd *NullUUID) Scan(value interface{}) (err error) {
	var s uuid.UUID
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if value == nil {
		*nd = NullUUID{Valid: false}
	} else {
		*nd = NullUUID{s, true}
	}

	return nil
}

func (nd NullUUID) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}

	return nd.UUID, nil // 修正: nd.UUID を返す
}
