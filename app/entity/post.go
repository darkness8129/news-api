package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Title   string
	Content string

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
