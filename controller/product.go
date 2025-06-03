package controller

import (
	"fmt"
	"net/http"
	config "store1/data"
	models "store1/model"

	"github.com/gin-gonic/gin"
)

// Tạo sản phẩm
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Product:", product)
	config.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

// Lấy tất cả sản phẩm
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

// Lấy 1 sản phẩm theo ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sản phẩm không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Cập nhật sản phẩm
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sản phẩm không tồn tại"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&product).Updates(updatedProduct)
	c.JSON(http.StatusOK, product)
}

// Xoá sản phẩm
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sản phẩm không tồn tại"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá sản phẩm thành công"})
}

func SearchProducts(keyword string) ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Where("name LIKE ?", keyword+"%").Find(&products).Error
	return products, err
}

// Lấy sản phẩm theo loại
func GetProductsByCategoryID(id string) ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Where("category_id = ?", id).Find(&products).Error
	return products, err
}

// Lọc sản phẩm theo giá với nhiều options
func FilterProductsByPrice(sortBy, priceRange string) ([]models.Product, error) {
	var products []models.Product
	query := config.DB

	// Áp dụng filter theo khoảng giá
	switch priceRange {
	case "under100k":
		query = query.Where("price < ?", 100000)
	case "100k-200k":
		query = query.Where("price >= ? AND price <= ?", 100000, 200000)
	case "over200k":
		query = query.Where("price > ?", 200000)
		// Không có filter nếu priceRange rỗng hoặc "all"
	}

	// Áp dụng sắp xếp
	switch sortBy {
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	case "name_asc":
		query = query.Order("name ASC")
	case "newest":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("id ASC") // Mặc định theo ID
	}

	err := query.Find(&products).Error
	return products, err
}

// Lọc sản phẩm theo giá với nhiều options có loại
func FilterProductsByCategoryAndPrice(categoryID, sortBy, priceRange string) ([]models.Product, error) {
	var products []models.Product
	query := config.DB.Where("category_id = ?", categoryID)

	// Áp dụng filter theo khoảng giá
	switch priceRange {
	case "under100k":
		query = query.Where("price < ?", 100000)
	case "100k-200k":
		query = query.Where("price >= ? AND price <= ?", 100000, 200000)
	case "over200k":
		query = query.Where("price > ?", 200000)
	}

	// Áp dụng sắp xếp
	switch sortBy {
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	case "name_asc":
		query = query.Order("name ASC")
	case "newest":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("id ASC")
	}

	err := query.Find(&products).Error
	return products, err
}
