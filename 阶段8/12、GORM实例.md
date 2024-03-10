```
go// models/user.go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string
    Age   uint
}

// dao/user_dao.go
package dao

import (
	"gorm.io/gorm"
	"your_project/models"
)

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// GetUserByID 根据ID获取用户
func GetUserByID(db *gorm.DB, userID uint) (models.User, error) {
	var user models.User
	err := db.First(&user, userID).Error
	return user, err
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}

// DeleteUser 删除用户
func DeleteUser(db *gorm.DB, userID uint) error {
	return db.Delete(&models.User{}, userID).Error
}

// GetAllUsers 获取所有用户
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}
```

以上是一个简单的 DAO（Data Access Object）层，用于执行数据库操作。创建一个处理 HTTP 请求的控制器或处理程序：

```
go// controllers/user_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"your_project/dao"
	"your_project/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dao.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUserByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := dao.GetUserByID(db, uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// 同样，你需要创建 UpdateUser、DeleteUser 和 GetAllUsers 的处理函数
```

