package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string     `json:"name"`
	Account     string     `json:"account"`
	Password    string     `json:"password"`
	ExpireAt    *time.Time `json:"expire_at"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
}
