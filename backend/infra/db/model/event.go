package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/iChemy/MyKnoQ/backend/domain"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID        uuid.UUID        `gorm:"primaryKey;type:char(36);"`
	Name      string           `gorm:"not null;"`
	StartAt   time.Time        `gorm:"not null;"`
	EndAt     time.Time        `gorm:"not null;"`
	EventType domain.EventType `gorm:"not null;type:smallInt;"`

	GroupID uuid.UUID `gorm:"type:char(36);not null;"`
	Group   Group     `gorm:"foreignKey:GroupID;"`

	CreatedByID uuid.UUID `gorm:"type:char(36);not null;"`
	CreatedBy   User      `gorm:"foreignKey:CreatedByID;"`

	Shareable sql.NullBool `gorm:""`
	RoomID    NullUUID     `gorm:"type:char(36);"`

	PrivateVenueName sql.NullString `gorm:"type:varchar(128);"`

	EventAdmins []User `gorm:"many2many:evnet_admin;"`

	Tags []Tag `gorm:"many2many:event_tag;"`

	Attendee []Participant `gorm:"foreignKey:EventID;"`
}

type Participant struct {
	UserID           uuid.UUID               `gorm:"primaryKey;type:char(36);"`
	EventID          uuid.UUID               `gorm:"primaryKey;type:char(36);"`
	ParticipantState domain.ParticipantState `gorm:"not null;type:smallInt;"`
}
