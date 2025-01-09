package domain

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	// Optional
	DeletedAt *time.Time
}
