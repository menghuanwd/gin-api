package models

type Product struct {
	Base
	Category   Category
	CategoryID uint
	Name       string  `json:"name" gorm:"index"`
	Price      float32 `json:"price"`
}
