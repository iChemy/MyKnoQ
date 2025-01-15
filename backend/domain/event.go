package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type EventType uint

const (
	EventTypeRoom EventType = iota
	EventTypePrivate
	EventTypeLimit
)

type Event interface {
	GetEventCore() EventCore
}

type EventCore struct {
	ID        uuid.UUID
	Name      string
	StartAt   time.Time
	EndAt     time.Time
	EventType EventType
	VenueName string

	CreatedBy   User
	EventAdmins []User

	Attendee []Participant
}

type RoomEvent struct {
	Model
	EventCore
	Shareable bool
}

func (re RoomEvent) GetEventCore() EventCore {
	return re.EventCore
}

type PrivateEvent struct {
	Model
	EventCore
}

func (pe PrivateEvent) GetEventCore() EventCore {
	return pe.EventCore
}

func (e *EventType) Scan(src interface{}) error {
	s := sql.NullByte{}
	if err := s.Scan(src); err != nil {
		return err
	}

	if s.Valid {
		newET := EventType(s.Byte)
		if newET >= EventTypeLimit {
			return fmt.Errorf("EventType(%d) must be less than %d", newET, EventTypeLimit)
		}

		*e = newET
	}

	return nil
}

func (e EventType) Value() (driver.Value, error) {
	return sql.NullByte{Byte: byte(e), Valid: true}.Value()
}

type ParticipantState uint

const (
	Unconfirmed ParticipantState = iota
	Attend
	Absent
	ParticipantStateLimit
)

func (p *ParticipantState) Scan(src interface{}) error {
	s := sql.NullByte{}
	if err := s.Scan(src); err != nil {
		return err
	}

	if s.Valid {
		newPS := ParticipantState(s.Byte)
		if newPS >= ParticipantStateLimit {
			return fmt.Errorf("ParticipantState(%d) must be less than %d", newPS, ParticipantStateLimit)
		}

		*p = newPS
	}

	return nil
}

func (p ParticipantState) Value() (driver.Value, error) {
	return sql.NullByte{Byte: byte(p), Valid: true}.Value()
}

type Participant struct {
	User             User
	ParticipantState ParticipantState
}
