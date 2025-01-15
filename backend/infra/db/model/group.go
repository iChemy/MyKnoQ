package model

import (
	"database/sql"

	"github.com/gofrs/uuid"
	"github.com/iChemy/MyKnoQ/backend/domain"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID        uuid.UUID        `gorm:"primaryKey;type:char(36);"`
	GroupType domain.GroupType `gorm:"type:smallInt;"`

	TraqID NullUUID `gorm:"type:char(36);"`

	JoinFree sql.NullBool `gorm:""`

	Name string `gorm:""`

	Members []Membership `gorm:"foreignKey:GroupID;"`
}

type Membership struct {
	UserID  uuid.UUID `gorm:"type:char(36);primaryKey;"`
	User    User      `gorm:"foreignKey:UserID;"`
	GroupID uuid.UUID `gorm:"primaryKey;type:char(36);"`
}
