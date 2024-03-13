### 模型定义

首先，定义一个结构体来表示数据库表。结构体中的字段对应于表中的列。

```go
package main  
  
import (  
	"gorm.io/gorm"  
)  
  
type User struct {  
	gorm.Model 
	Name       string  
	Age        int  
}
```

在这个例子中，`User`结构体包含了`gorm.Model`，它会自动为模型添加`ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段。你也可以选择不包含`gorm.Model`，并自行定义主键和其他字段。这里有前面笔记提及到`gorm.Model`，了解到它具有一定局限性，但会很方便的帮助我们创建一些常见的结构体。所以说有利有弊，我个人认为这个类似与于c语言中的stl库。虽然是远远比不上stl库的复杂程度。

### 连接数据库

在执行任何CRUD操作之前，需要先连接到数据库。

```go
import (  
	"gorm.io/driver/mysql"  
	"gorm.io/gorm"  
)  
  
func main() {  
	dsn := "root:Hjh994600!@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  
	if err != nil {  
		panic("failed to connect database")  
	}  
	// 自动迁移
	db.AutoMigrate(&User{})  
}
```

### CRUD操作

#### 创建

```go
func createUser(db *gorm.DB) {  
	user := User{Name: "John Doe", Age: 30}  
	result := db.Create(&user) // 返回创建记录的结果  
	if result.Error != nil {  
		// 处理错误  
	}  
}
```

#### 读取

```go
func readUser(db *gorm.DB, id uint) {  
	var user User  
	result := db.First(&user, id) // 根据主键读取  
	if result.Error != nil {  
		// 处理错误  
	}  
	// 打印用户信息  
	fmt.Println(user.Name, user.Age)  
}
```

#### 更新

```go
func updateUser(db *gorm.DB, id uint) {  
	user := User{Name: "Jane Doe", Age: 35}  
	result := db.Model(&user).Where("id = ?", id).Updates(User{Name: "Updated Name", Age: 0})  
	if result.Error != nil {  
		// 处理错误  
	}  
}
```

#### 删除

```go
func deleteUser(db *gorm.DB, id uint) {  
	user := User{}  
	result := db.Delete(&user, id) // 根据主键删除  
	if result.Error != nil {  
		// 处理错误  
	}  
}
```

### 查询

#### 简单查询

```go
func simpleQuery(db *gorm.DB) {  
	var users []User  
	result := db.Find(&users) // 查询所有用户  
	if result.Error != nil {  
		// 处理错误  
	}  
	// 遍历用户列表  
	for _, user := range users {  
		fmt.Println(user.Name, user.Age)  
	}  
}
```

#### 条件查询

```go
func conditionalQuery(db *gorm.DB) {  
	var users []User  
	result := db.Where("name = ?", "John Doe").Find(&users) // 根据条件查询用户  
	if result.Error != nil {  
		// 处理错误  
	}  
	// 遍历用户列表  
	for _, user := range users {  
		fmt.Println(user.Name, user.Age)  
	}  
}
```

#### 排序

```go
func sortedQuery(db *gorm.DB) {  
	var users []User  
	result := db.Order("age desc").Find(&users) // 按年龄降序排序  
	if result.Error != nil {  
		// 处理错误  
	}  
	// 遍历用户列表  
	for _, user := range users {  
		fmt.Println(user.Name, user.Age)  
	}  
}
```

#### 分页

在GORM中，你可以使用`Limit`和`Offset`方法进行分页查询。下面是一个分页查询的示例：

```go
func paginatedQuery(db *gorm.DB, pageNumber, pageSize int) {  
	var users []User  
	var total int64 // 总记录数  
  
	// 使用Offset和Limit进行分页  
	result := db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Find(&users).Count(&total)  
	if result.Error != nil {  
		// 处理错误  
	}  
  
	// 打印当前页的用户列表  
	for _, user := range users {  
		fmt.Println(user.Name, user.Age)  
	}  
  
	// 打印总记录数  
	fmt.Printf("Total records: %d\n", total)  
}
```

在这个例子中，`pageNumber`表示页码（通常从1开始），`pageSize`表示每页的记录数。`Offset`方法用于跳过前面的记录，而`Limit`方法则限制返回的记录数。`Count`方法用于获取总记录数，通常用于显示给用户知道总共有多少条记录。

### 错误处理

在进行CRUD和查询操作时，你可能会遇到错误。GORM的`result`对象包含一个`Error`字段，你可以检查这个字段来判断操作是否成功。

```go
result := db.Create(&user)  
if result.Error != nil {  
    // 处理错误，例如打印错误信息或返回错误给调用者  
    fmt.Println("Error creating user:", result.Error)  
}
```

### 示例：完整的CRUD和查询操作

下面是一个包含完整CRUD和查询操作的简单示例：

```go
package main  
  
import (  
	"fmt"  
	"gorm.io/driver/mysql"  
	"gorm.io/gorm"  
)  
  
type User struct {  
	gorm.Model  
	Name string  
	Age  int  
}  
  
func main() {  
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  
	if err != nil {  
		panic("failed to connect database")  
	}  
  
	// 自动迁移schema  
	db.AutoMigrate(&User{})  
  
	// 创建用户  
	createUser(db)  
  
	// 读取用户  
	readUser(db, 1)  
  
	// 更新用户  
	updateUser(db, 1)  
  
	// 再次读取用户以验证更新  
	readUser(db, 1)  
  
	// 删除用户  
	deleteUser(db, 1)  
  
	// 简单查询所有用户  
	simpleQuery(db)  
  
	// 条件查询用户  
	conditionalQuery(db)  
  
	// 排序查询用户  
	sortedQuery(db)  
  
	// 分页查询用户  
	paginatedQuery(db, 1, 2) // 假设每页2条记录，查询第1页  
}  
  
// ... (上面定义的CRUD和查询函数)  
  
// 注意：上面的函数需要按照它们在main函数中调用的顺序进行定义。
```

在上面的`main`函数中，我们演示了如何使用GORM进行数据库连接、模型迁移、创建、读取、更新、删除以及各种查询操作。每个操作都封装在了一个函数中，以便在`main`函数中按顺序调用。请确保你已经正确设置了数据库连接字符串，并且数据库服务器正在运行。





**上文对dsn产生疑问，文末进行简单了解**

**数据源**：DSN（Data Source Name）是ODBC（开放数据库连接）中定义的一个术语，它为一个确定的数据库和必须用到的ODBC驱动程序提供了标识。每个ODBC驱动程序都定义了该驱动程序支持的一个数据库创建DSN所需的信息。在安装ODBC驱动程序并创建一个数据库之后，必须创建一个DSN。DSN文件（数据源名）主要用于**存储数据库连接信息**。

个人认为这是一个类似内存但是临时储存的地方，更像是栅。