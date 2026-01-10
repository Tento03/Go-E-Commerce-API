package product

type UpdateProduct struct {
	Title       string `json:"title" binding:"required,min=5"`
	Description string `json:"description"`
}
