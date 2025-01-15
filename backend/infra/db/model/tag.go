package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primaryKey;type:char(36);"`
	Name string    `gorm:"not null;uniqueIndex;"`
}
