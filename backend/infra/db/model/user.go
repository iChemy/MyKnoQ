package model

import (
	"github.com/gofrs/uuid"
	"github.com/iChemy/MyKnoQ/backend/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:char(36);"`
	Name       string    `gorm:"uniqueIndex;not null;"`
	Privileged bool      `gorm:"not null;"`

	DisplayName   string               `gorm:""`
	TraqUserState domain.TraqUserState `gorm:"type:smallInt;"`
}
