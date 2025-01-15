package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/gofrs/uuid"
)

type TraqUserState uint

const (
	TraqStateDeactivated TraqUserState = iota
	// ユーザーアカウント状態: 有効
	TraqStateActive
	// ユーザーアカウント状態: 一時停止
	TraqStateSuspended
	TraqStateLimit
)

type User struct {
	ID            uuid.UUID
	Name          string
	DisplayName   string
	TraqUserState TraqUserState
}

func (t *TraqUserState) Scan(src interface{}) error {
	s := sql.NullByte{}
	if err := s.Scan(src); err != nil {
		return err
	}

	if s.Valid {
		newTS := TraqUserState(s.Byte)
		if newTS >= TraqStateLimit {
			return fmt.Errorf("TraqUserState(%d) must be less than %d", newTS, TraqStateLimit)
		}

		*t = newTS
	}

	return nil
}

func (t TraqUserState) Value() (driver.Value, error) {
	return sql.NullByte{Byte: byte(t), Valid: true}.Value()
}
