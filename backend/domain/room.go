package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Room struct {
	ID      uuid.UUID
	Name    string
	StartAt time.Time
	EndAt   time.Time

	Events []RoomEvent
}
