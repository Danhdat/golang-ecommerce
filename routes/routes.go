package routes

import (
	controllers "store1/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	index := router.Group("/")
	{
		index.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// Danh sách sản phẩm
	router.GET("/products", func(c *gin.Context) {
		products, err := controllers.GetProducts()
		if err != nil {
			c.JSON(500, gin.H{"error": "Không thể lấy sản phẩm"})
			return
		}

		//Lấy tất cả danh mục
		allCategories, _ := controllers.GetCategories()

		c.HTML(200, "shop.html", gin.H{
			"Products":   products,
			"Categories": allCategories,
			"BaseURL":    "/",
		})
	})

	router.GET("/categories/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Lấy danh mục theo id
		category, err := controllers.GetCategoryByID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Không tìm thấy danh mục"})
			return
		}
		categoryName := category[0].Name

		// Lấy sản phẩm theo category
		products, err := controllers.GetProductsByCategoryID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Lỗi khi lấy sản phẩm"})
			return
		}

		//Lấy tất cả danh mục
		allCategories, _ := controllers.GetCategories()

		c.HTML(200, "shop.html", gin.H{
			"Products":       products,
			"Categories":     allCategories,
			"ActiveCategory": categoryName,
			"CategoryID":     id,
			"BaseURL":        "/",
		})
	})

	router.GET("/products/search", func(c *gin.Context) {
		q := c.Query("q") // từ khoá tìm kiếm

		products, err := controllers.SearchProducts(q)
		if err != nil {
			c.HTML(500, "error.html", gin.H{"error": "Lỗi khi tìm sản phẩm"})
			return
		}

		allCategories, _ := controllers.GetCategories()

		c.HTML(200, "shop.html", gin.H{
			"Products":      products,
			"Categories":    allCategories,
			"SearchKeyword": q,
			"BaseURL":       "/",
		})
	})

	// Lọc sản phẩm theo giá
	router.GET("/products/filter", func(c *gin.Context) {
		sortBy := c.DefaultQuery("sort", "")      // price_asc, price_desc, name_asc, newest
		priceRange := c.DefaultQuery("price", "") // under100k, 100k-200k, over200k
		categoryID := c.Query("category")         // Optional: filter by category too

		var products interface{}
		var err error
		var activeCategory string

		// Nếu có category filter
		if categoryID != "" {
			products, err = controllers.FilterProductsByCategoryAndPrice(categoryID, sortBy, priceRange)
			if category, catErr := controllers.GetCategoryByID(categoryID); catErr == nil && len(category) > 0 {
				activeCategory = category[0].Name
			}

		} else {
			products, err = controllers.FilterProductsByPrice(sortBy, priceRange)
		}

		if err != nil {
			c.JSON(500, gin.H{"error": "Lỗi khi lọc sản phẩm"})
			return
		}

		// Lấy tất cả categories cho sidebar
		allCategories, _ := controllers.GetCategories()

		c.HTML(200, "shop.html", gin.H{
			"Products":       products,
			"Categories":     allCategories,
			"ActiveCategory": activeCategory,
			"CategoryID":     categoryID,
			"BaseURL":        "/",
		})
	})

	cartGroup := router.Group("/api/cart")
	{
		// GET /api/cart - Lấy giỏ hàng hiện tại
		cartGroup.GET("", controllers.GetCart)
	}

}
