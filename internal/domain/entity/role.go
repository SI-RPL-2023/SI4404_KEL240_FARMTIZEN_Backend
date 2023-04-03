package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID   string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}
