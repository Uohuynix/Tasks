**1. 模型定义：**

假设我们有一个名为 `User` 的结构体，用于表示数据库中的用户表：

```
gopackage models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string
    Age   uint
}
```

在这个例子中，`User` 结构体包含了 `Name`、`Email` 和 `Age` 字段，同时使用了 `gorm.Model` 结构体，它包含了一些公共字段，如 `ID`、`CreatedAt`、`UpdatedAt` 和 `DeletedAt`，这些字段用于跟踪记录的创建、更新和删除时间。

**2. CRUD 操作：**

- **创建（Create）数据：**

```
gofunc CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}
```

- **读取（Read）数据：**

```
gofunc GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
    var user models.User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
```

- **更新（Update）数据：**

```
gofunc UpdateUser(db *gorm.DB, user *models.User, newData map[string]interface{}) error {
    return db.Model(user).Updates(newData).Error
}
```

- **删除（Delete）数据：**

```
gofunc DeleteUser(db *gorm.DB, user *models.User) error {
    return db.Delete(user).Error
}
```

**3. 查询：**

- **简单查询：**

```
gofunc GetAllUsers(db *gorm.DB) ([]models.User, error) {
    var users []models.User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
```

- **条件查询：**

```
gofunc GetUsersByAge(db *gorm.DB, age uint) ([]models.User, error) {
    var users []models.User
    if err := db.Where("age = ?", age).Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
```

- **排序：**

```
gofunc GetUsersSortedByName(db *gorm.DB) ([]models.User, error) {
    var users []models.User
    if err := db.Order("name").Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
```

- **分页：**

```
gofunc GetUsersPaginated(db *gorm.DB, page, pageSize int) ([]models.User, error) {
    var users []models.User
    if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
```