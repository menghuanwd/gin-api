package vos

type Category struct {
	Base
	Name   string `json:"name"`
	Image  string `json:"image"`
	ShopID uint   `json:"shop_id,omitempty"`
}
