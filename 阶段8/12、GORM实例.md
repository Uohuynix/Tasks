```
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Product 模型
type Product struct {
	ID    uint
	Code  string
	Price uint
}

func main() {
	// 连接到 MySQL 数据库
	dsn := "root:Hjh994600!@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移模式，以确保模型对应的表存在
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	// 创建记录
	product := Product{Code: "Laptop", Price: 1000}
	result := db.Create(&product)
	if result.Error != nil {
		log.Fatalf("Failed to create product: %v", result.Error)
	}

	// 读取记录
	var retrievedProduct Product
	result = db.First(&retrievedProduct, product.ID)
	if result.Error != nil {
		log.Fatalf("Failed to retrieve product: %v", result.Error)
	}
	log.Printf("Retrieved product: %+v", retrievedProduct)

	// 更新记录
	retrievedProduct.Price = 1200
	result = db.Save(&retrievedProduct)
	if result.Error != nil {
		log.Fatalf("Failed to update product: %v", result.Error)
	}

	// 删除记录
	result = db.Delete(&retrievedProduct, retrievedProduct.ID)
	if result.Error != nil {
		log.Fatalf("Failed to delete product: %v", result.Error)
	}
}

```

