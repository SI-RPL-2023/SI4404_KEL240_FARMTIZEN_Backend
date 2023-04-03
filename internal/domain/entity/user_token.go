package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	ID     string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID string `json:"user_id" gorm:"type:uuid;not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Type   string `json:"type" gorm:"type:varchar(255);not null"`
	Token  string `json:"token" gorm:"type:varchar(255);not null"`
}

func (ut *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	ut.ID = uuid.New().String()
	return
}
