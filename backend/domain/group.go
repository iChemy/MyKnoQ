package domain

import (
	"github.com/google/uuid"
)

type GroupType int

const (
	TypeKnoQGroup GroupType = iota
	TypeTraQGroup
)

type Group interface {
	GetGroupType() GroupType
}

type GroupCore struct {
	ID          uuid.UUID
	Name        string
	Description string
	Members     []User
	Admins      []User
	Model
}

type KnoQGroup struct {
	GroupCore
	JoinFree bool
}

func (g KnoQGroup) GetGroupType() GroupType {
	return TypeKnoQGroup
}

type TraQGroup struct {
	GroupCore
	TraQGroupID uuid.UUID
}

func (g TraQGroup) GetGroupType() GroupType {
	return TypeTraQGroup
}
