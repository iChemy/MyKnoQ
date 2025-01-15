package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/gofrs/uuid"
)

type GroupType uint

const (
	GroupTypeTraq GroupType = iota
	GroupTypeKnoq
	GroupTypeLimit
)

type GroupCore struct {
	Model
	ID        uuid.UUID
	GroupType GroupType
	Name      string
	Members   []Membership
}

type MembershipType uint

const (
	MembershipAdmin MembershipType = iota
	MembershipMember
	MembershipLimit
)

type Membership struct {
	User           User
	MembershipType MembershipType
}

type Group interface {
	GetGroupCore() GroupCore
}

type TraqGroup struct {
	GroupCore
	TraqID uuid.UUID
}

func (tg TraqGroup) GetGroupCore() GroupCore {
	return tg.GroupCore
}

type KnoqGroup struct {
	GroupCore
	JoinFree bool
}

func (kg KnoqGroup) GetGroupCore() GroupCore {
	return kg.GroupCore
}

func (g *GroupType) Scan(src interface{}) error {
	s := sql.NullByte{}
	if err := s.Scan(src); err != nil {
		return err
	}

	if s.Valid {
		newGT := GroupType(s.Byte)
		if newGT >= GroupTypeLimit {
			return fmt.Errorf("GroupType(%d) must be less than %d", newGT, GroupTypeLimit)
		}

		*g = newGT
	}

	return nil
}

func (g GroupType) Value() (driver.Value, error) {
	return sql.NullByte{Byte: byte(g), Valid: true}.Value()
}

func (m *MembershipType) Scan(src interface{}) error {
	s := sql.NullByte{}
	if err := s.Scan(src); err != nil {
		return err
	}

	if s.Valid {
		newMT := MembershipType(s.Byte)
		if newMT >= MembershipLimit {
			return fmt.Errorf("MembershipType(%d) must be less than %d", newMT, MembershipLimit)
		}

		*m = newMT
	}

	return nil
}

func (m MembershipType) Value() (driver.Value, error) {
	return sql.NullByte{Byte: byte(m), Valid: true}.Value()
}
