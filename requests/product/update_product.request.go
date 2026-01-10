package product

type UpdateProduct struct {
	Title       string `form:"title" binding:"required,min=5"`
	Description string `form:"description"`
	Price       string `form:"price" binding:"required,min=3"`
	Path        string `form:"path" binding:"required"`
}
