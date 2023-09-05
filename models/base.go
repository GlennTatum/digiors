package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

func (v *Model) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.NewString()
	return nil
}
