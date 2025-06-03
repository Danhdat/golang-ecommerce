package main

import (
	config "store1/data"
	models "store1/model"
	routes "store1/routes"

	"fmt"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Kết nối DB
	config.ConnectDB()

	// Auto migrate các model
	config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.CartItem{},
		&models.Payment{},
	)
	r.SetFuncMap(template.FuncMap{
		"formatPrice": formatPrice,
	})
	// Dùng để phục vụ tĩnh (CSS, JS, IMG...)
	r.Static("/assets", "./templates/assets")

	// Dùng để load file HTML
	r.LoadHTMLGlob("templates/*.html")
	// Đăng ký route
	routes.RegisterRoutes(r)
	// Chạy server
	r.Run(":8080")
}

// format hiển thị giá
func formatPrice(price float64) string {
	s := fmt.Sprintf("%.0f", price) // loại bỏ phần thập phân
	n := len(s)
	if n <= 3 {
		return s + " đ"
	}

	var result []string
	for i := n; i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}

	return strings.Join(result, ",") + " đ"
}
