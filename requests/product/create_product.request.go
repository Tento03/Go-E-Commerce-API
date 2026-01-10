package product

type CreateProduct struct {
	Title       string `form:"title" binding:"required,min=5"`
	Description string `form:"description"`
}
