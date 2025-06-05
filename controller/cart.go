package controller

import (
	"net/http"

	config "store1/data"
	models "store1/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCart(c *gin.Context) {
	// Lấy sessionID từ cookie (hoặc tạo mới nếu chưa có)
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		// Tạo session ID mới nếu chưa có
		sessionID = uuid.New().String()
		c.SetCookie("session_id", sessionID, 3600*24*7, "/", "", false, true)
	}

	// Kiểm tra nếu có user đăng nhập
	userID, loggedIn := c.Get("userID")

	var cart models.Cart
	var query interface{}

	if loggedIn {
		// Nếu đã đăng nhập: tìm cart bằng userID
		query = userID
	} else {
		// Nếu không đăng nhập: tìm cart bằng sessionID
		query = sessionID
	}

	// Tìm hoặc tạo giỏ hàng
	result := config.DB.Where("user_id = ? OR session_id = ?", query, sessionID).
		Preload("Items").
		Preload("Items.Product").
		First(&cart)

	if result.Error != nil {
		// Tạo giỏ hàng mới
		cart = models.Cart{
			SessionID: sessionID,
		}
		if loggedIn {
			cart.UserID = userID.(uint)
		}
		config.DB.Create(&cart)
	}

	// Tính tổng tiền
	var total float64
	for _, item := range cart.Items {
		total += item.Product.Price * float64(item.Quantity)
	}

	c.JSON(http.StatusOK, gin.H{
		"cart":  cart,
		"total": total,
	})
}
