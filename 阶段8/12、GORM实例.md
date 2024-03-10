```
gopackage main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义一个模型
type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB

// 初始化数据库连接
func InitDB() {
	// 连接 SQLite 数据库，你可以根据需要选择其他数据库
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}

	// 迁移模型，创建表
	database.AutoMigrate(&User{})

	db = database
}

func main() {
	// 初始化数据库连接
	InitDB()

	// 关闭数据库连接
	defer db.Close()

	// 在这里你可以执行各种数据库操作，例如创建、查询、更新、删除等

	// 创建用户
	user := User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)

	// 查询所有用户
	var users []User
	db.Find(&users)
	fmt.Println("所有用户：", users)
}
```

