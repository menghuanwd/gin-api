package models

import "time"

type Editor struct {
	Base
	Shops       []Shop    `json:"-"`
	Name        string    `json:"name" gorm:"index"`
	Mobile      string    `json:"mobile" gorm:"unique"`
	Email       string    `json:"email"  gorm:"unique"`
	LastLoginAt time.Time `json:"last_login_at"`
	AuthToken   string    `json:"auth_token"`
}
