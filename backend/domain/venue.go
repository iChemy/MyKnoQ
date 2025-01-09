package domain

import (
	"time"

	"github.com/google/uuid"
)

type VenueType int

const (
	TypePrivate VenueType = iota
	TypeRoom
)

type Venue interface {
	GetVenueType() Venue
}

type VenueCore struct {
	Name    string
	StartAt time.Time
	EndAt   time.Time
}

type PrivateVenue struct {
	VenueCore
}

func (v PrivateVenue) GetGroupType() VenueType {
	return TypePrivate
}

type RoomVenue struct {
	VenueCore
	Shareable bool
	RoomID    uuid.UUID
}

func (v RoomVenue) GetGroupType() VenueType {
	return TypeRoom
}
