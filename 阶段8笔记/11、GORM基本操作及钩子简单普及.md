摘要：本片笔记将对gorm的一些基本操作进行描述，过程中遇到许多问题将在本文结尾提及。

------

GORM是一个流行的Go语言ORM（对象关系映射）库，它使得开发者能够使用Go的结构体来操作关系型数据库，如MySQL。以下是使用GORM进行基本操作的一些步骤和示例：

### 1. 安装GORM

首先，你需要安装GORM库。可以使用`go get`命令来获取：

```bash
go get -u gorm.io/gorm  
go get -u gorm.io/driver/mysql
```

### 2. 连接数据库

连接数据库是使用GORM的第一步。需要提供数据库驱动名称（如MySQL）和数据源名称（DSN），DSN中包含了连接数据库所需的所有信息。

```go
import (  
    "gorm.io/gorm"  
    "gorm.io/driver/mysql"  
)  
  
func main() {  
    dsn := "root:Hjh994600!@tcp(localhost:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"  
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  
    if err != nil {  
        panic("failed to connect database")  
    }  
    // 自动迁移（目前了解不够）
    db.AutoMigrate(&User{})  
    //进行其他操作
}
```

### 3. 定义模型

在Go中，你通常使用结构体来表示数据库中的表。GORM会根据结构体的字段自动映射到数据库的列。

```go
type User struct {  
    gorm.Model // 包含ID, CreatedAt, UpdatedAt, DeletedAt字段  
    Name       string  
    Age        int  
    Email      string  
}
```

### 4. 创建记录

使用`Create`方法向数据库中插入数据。

```go
user := User{Name: "John Doe", Age: 30, Email: "johndoe@example.com"}  
result := db.Create(&user)  
if result.Error != nil {  
    // 处理错误  
}
```

### 5. 查询记录

GORM提供了多种查询方法，如`First`、`Find`、`Where`等。

```go
var user User  
db.First(&user, 1) // 查询ID为1的用户  
  
var users []User  
db.Find(&users) // 查询所有用户  
  
// 使用条件查询  
db.Where("name = ?", "John Doe").Find(&users)
```

### 6. 更新记录

使用`Save`或`Updates`方法更新数据库中的记录。

```go
// 根据主键更新记录  
db.Save(&user)  
  
// 更新特定字段  
db.Model(&user).Updates(User{Age: 31})
```

### 7. 删除记录

使用`Delete`方法删除数据库中的记录。

```go
db.Delete(&user, 1) // 删除ID为1的用户
```

### 8. 高级操作（尚未深入学习）

GORM还支持事务操作、关联查询、钩子、预加载等高级功能。例如，使用事务可以确保多个操作要么全部成功，要么全部失败。

```go
tx := db.Begin()  
if err := tx.Create(&user).Error; err != nil {  
    tx.Rollback() // 如果出现错误，回滚事务  
} else {  
    tx.Commit() // 提交事务  
}
```

### 注意事项

- 不要在核心结构体中加入非表中的数据，如一些计算的中间值，以免引起二义性。
- `gorm.Model` 可以提升编码效率（会减少重复编码），但会限制数据库表中字段的定义，因此在使用时要谨慎。
- 在查询和更新接口里，推荐的使用方法是采用核心结构体加上一个字段数组的方式，前者保存具体的数据、也实现了结构体复用，后者则选择生效的字段。



### 钩子简单了解：

GORM 的钩子是一种在模型生命周期中自动执行特定函数的功能。这些钩子函数允许你在模型的创建、保存、更新、删除等操作中执行自定义逻辑。

GORM 提供了以下钩子函数（等待继续学习）：

1. **BeforeCreate**：在创建记录到数据库之前调用。
2. **AfterCreate**：在创建记录到数据库之后调用。
3. **BeforeSave**：在创建或更新记录到数据库之前调用。
4. **AfterSave**：在创建或更新记录到数据库之后调用。
5. **BeforeUpdate**：在更新记录到数据库之前调用。
6. **AfterUpdate**：在更新记录到数据库之后调用。
7. **BeforeDelete**：在删除记录之前调用。
8. **AfterDelete**：在删除记录之后调用。
9. **BeforeFind**：在查询数据库之前调用。
10. **AfterFind**：在查询数据库之后调用。

这些钩子函数可以通过在模型结构体中定义相应的方法来实现。例如，每次创建用户时**自动设置其创建时间**，你可以这样做：

```go
type User struct {  
    gorm.Model // 包含ID, CreatedAt, UpdatedAt, DeletedAt字段  
    Name       string  
    // 其他操作
}  
  
// 实现BeforeCreate钩子  
func (u *User) BeforeCreate(db *gorm.DB) error {  
    u.CreatedAt = time.Now() // 设置创建时间为当前时间  有在猜拳游戏中使用到
    return nil  
}
```

在上面的例子中，`BeforeCreate` 方法会在每次创建 `User` 记录之前被自动调用，并设置 `CreatedAt` 字段的值为当前时间。

多个模型中**共享相同的钩子逻辑**，你可以使用**回调函数**而不是在模型中定义方法。例如：

```go
// 定义一个通用的创建前钩子函数  
func commonBeforeCreate(scope *gorm.Scope) error {  
    if !scope.HasError() {  
        scope.SetColumn("CreatedAt", time.Now())  
    }  
    return nil  
}  
  
// 在初始化数据库连接时注册钩子  
db.Callback().Create().Before("gorm:create").Register("common_before_create", commonBeforeCreate)
```

在这个例子中，`commonBeforeCreate` 函数是一个通用的创建前钩子，它会在任何模型的创建操作之前设置 `CreatedAt` 字段。然后，我们通过 `db.Callback().Create().Before("gorm:create").Register` 方法将这个钩子注册到 GORM 的回调链中。

Tips：GORM 的钩子函数通常接受一个指向 `gorm.DB`（对于方法钩子）或 `gorm.Scope`（对于回调函数钩子）的指针作为参数，并返回一个错误（如果有的话）。



**所遇问题及解决方式：**

首先是数据库的下载及连接问题，首次下载无法正常获取解析包，正常下载后也没办法在import里打开，这个问题搜索了许多方法去解决，终于在csdn里找到了一篇文章，然后在实验室里学长的指教下，重新创建go mod并清空历史go mod，在此过程中遇到的另一个问题是文件存放的位置不对，错误的把项目文件存放在了src包里（前期是存放在bin里，首次安装时遇到了问题，根据文章更改了项目位置，但是并没有更改正确），后续在独立于goland文件夹外创建一个新的文件夹，重新配置了环境并且安装

```bash
go get -u gorm.io/gorm  
go get -u gorm.io/driver/mysql
```

后续`import`便能正确打开并响应语句

```go
 "gorm.io/gorm"  
 "gorm.io/driver/mysql"  
```