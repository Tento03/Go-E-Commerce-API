package product

import (
	repositories "ecommerce-api/repositories/product"
	requests "ecommerce-api/requests/product"
	services "ecommerce-api/services/product"
	"fmt"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindAll(c *gin.Context) {
	product, err := repositories.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": product})
}

func FindById(c *gin.Context) {
	userId := c.Param("id")
	product, err := repositories.FindById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": product})
}

func FindByTitle(c *gin.Context) {
	userId := c.Param("id")
	product, err := repositories.FindById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	title := product.Title
	productTitle, err := repositories.FindByTitle(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": productTitle})
}

func CreateProduct(c *gin.Context) {
	var req requests.CreateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	if file.Size > 5<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file size max 5mb"})
		return
	}

	ext := strings.ToLower(filepath.Base(file.Filename))
	allowedType := map[string][]string{
		"jpg": {".jpg", ".jpeg"},
		"png": {".png"},
	}

	if !slices.Contains(allowedType[req.Type], ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file type not allowed"})
		return
	}

	productId := uuid.NewString()
	filename := fmt.Sprintf(
		"%s_%d%s",
		productId,
		time.Now().UnixNano(),
		ext,
	)
	path := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save uploaded file"})
		return
	}

	product, err := services.CreateProduct(productId, req.Title, req.Description, req.Price, req.Type, req.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": product})
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	product, err := services.GetById(productId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var req requests.UpdateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, _ := c.FormFile("file")
	if file != nil {
		if file.Size > 5<<20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file size max 5mb"})
			return
		}

		ext := strings.ToLower(filepath.Base(file.Filename))
		allowedType := map[string][]string{
			"jpg": {".jpg", ".jpeg"},
			"png": {".png"},
		}

		if !slices.Contains(allowedType[req.Type], ext) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file type not allowed"})
			return
		}

		productId := product.ProductID
		filename := fmt.Sprintf(
			"%s_%d%s",
			productId,
			time.Now().UnixNano(),
			ext,
		)
		path := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save uploaded file"})
			return
		}

		products, err := services.UpdateProduct(productId, req.Title, req.Description, req.Price, req.Type, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success update data", "data": products})
	}
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	err := services.DeleteProduct(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
