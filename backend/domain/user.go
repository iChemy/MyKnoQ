package domain

import (
	"github.com/google/uuid"
)

type UserAccountState int

const (
	AccountDead UserAccountState = iota
	AccountAlive
	AccountFreezed
)

type User struct {
	ID          uuid.UUID
	Name        string
	DisplayName string
	Icon        string
	Privileged  bool
	State       UserAccountState
}
