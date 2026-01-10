package product

type UpdateProductRequest struct {
	Title       string `form:"title" binding:"required,min=5"`
	Description string `form:"description"`
	Price       string `form:"price" binding:"required,min=3"`
	Type        string `form:"type" binding:"required,oneof=jpg png"`
	Path        string `form:"path" binding:"required"`
}
