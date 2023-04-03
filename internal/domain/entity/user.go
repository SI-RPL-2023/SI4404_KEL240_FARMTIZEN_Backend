package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName    string `json:"full_name" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Address     string `json:"address" gorm:"type:varchar(255);not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
