package params

type Category struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image"`
}
