package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`

	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
