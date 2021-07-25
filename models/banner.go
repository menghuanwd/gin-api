package models

type Banner struct {
	Base
	Shop   Shop
	ShopID uint
	Name   string `json:"name"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}
