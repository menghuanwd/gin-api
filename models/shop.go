package models

type Shop struct {
	Base
	Editor     Editor
	Category   []Category
	Users      []User
	EditorID   uint
	Name       string `json:"name" gorm:"index"`
	Status     string `json:"status"`
	Address    string `json:"address"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}
