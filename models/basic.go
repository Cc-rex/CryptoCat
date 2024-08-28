package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"`
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`
	UpdatedAt time.Time `json:"updated_at" structs:"-"`
}
