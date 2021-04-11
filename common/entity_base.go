package common

import "time"

type IDBase struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type UUIDBase struct {
	ID string `gorm:"primaryKey;type:VARCHAR(36)" json:"id"`
}

type TimeBase struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SoftDeleteBase struct {
	DeletedAt time.Time
}
