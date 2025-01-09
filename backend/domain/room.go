package domain

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID      uuid.UUID
	Name    string
	StartAt time.Time
	EndAt   time.Time
	Model

	CreatedBy User
}
