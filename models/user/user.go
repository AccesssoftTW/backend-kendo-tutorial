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
	ExpiredAt   *time.Time `json:"expired_at"`
	PhoneNumber string     `json:"phone_number"`
}
