package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Cấu hình kết nối MySQL
	dsn := "root@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("⚠️ Không thể kết nối đến database:", err)
	}

	fmt.Println("Kết nối database thành công")
	DB = db
}
