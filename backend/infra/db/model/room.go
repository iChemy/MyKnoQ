package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:char(36);"`
	Name        string    `gorm:"not null;"`
	StartAt     time.Time `gorm:"not null;"`
	EndAt       time.Time `gorm:"not null;"`
	CreatedByID uuid.UUID `gorm:"not null;type:char(36);"`
	CreatedBy   User
}
