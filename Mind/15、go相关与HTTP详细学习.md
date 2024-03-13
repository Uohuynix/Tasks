[TOC]

Go语言中的HTTP包提供了用于创建Web服务器和客户端的功能。	

### 1. 创建HTTP服务器

#### a. 创建基本的HTTP服务器

你可以使用包来创建一个简单的HTTP服务器，示例代码如下：`http`

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

在这个例子中，我们创建了一个基本的HTTP服务器。函数是一个处理器函数，它接收一个作为参数，可以用来向客户端发送HTTP响应；另外，还接收一个参数，用来获取客户端请求的信息。`handler``http.ResponseWriter``*http.Request`

#### b. 路由和处理器函数

```go
http`包中的函数可以用来指定不同路径下的处理器函数。比如：`HandleFunc
func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/about", aboutHandler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About Page")
}
```

在这个例子中，路径下的请求会被函数处理，路径下的请求会被函数处理。`/``handler``/about``aboutHandler`

- w：回写给客户端的数据
- req：读取到的客户端请求数据
- addr：服务器的ip:port
- 回调函数，通常传nil表示默认

### 2. HTTP 客户端

Go语言的包也提供了创建HTTP客户端的功能，可以用来向其他服务器发送HTTP请求。`http`

#### a. 发送 GET 请求

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(body))
}
```

在这个例子中，我们使用函数向指定的URL发送GET请求，并读取响应体中的内容。`http.Get`

#### b. 发送 POST 请求

```go
package main

import (
    "fmt"
    "bytes"
    "net/http"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/posts"
    jsonStr := []byte(`{"title":"Test Post","body":"This is a test post","userId":1}`)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
}
```

在这个例子中，我们使用函数创建了一个POST请求，并向指定的URL发送JSON数据。`http.NewRequest`

- 参数：待请求的url。**注意：必须带有http://开头**
- 返回值：服务器回复的应答包
- **Body是一个指针类型 – read，close**

### 3. 处理表单数据

#### a. 解析表单数据

```go
package main

import (
    "fmt"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("Form data:", r.Form)
    fmt.Println("Post data:", r.PostForm)
}

func main() {
    http.HandleFunc("/form", formHandler)
    http.ListenAndServe(":8080", nil)
}
```

在这个例子中，我们定义了一个处理表单数据的处理器函数。通过调用函数，我们可以解析HTTP请求中的表单数据，并将其存储在和中。`formHandler``r.ParseForm()``r.Form``r.PostForm`

### 4. 中间件

Go语言中的包还支持中间件，可以用来在请求到达处理器函数之前或之后执行一些逻辑。`http`

```go
package main

import (
    "fmt"
    "net/http"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Executing middleware...")
        next.ServeHTTP(w, r)
        fmt.Println("Middleware finished.")
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", middleware(handler))
    http.ListenAndServe(":8080", nil)
}
```

在这个例子中，我们定义了一个简单的中间件函数，它会在请求到达函数之前执行，并在之后输出一些日志信息。`middleware``handler`





### 5、对于一个HTTP请求可以细分为四个部分：

#### 	1、请求行：包含三部分

​			格式： 方法 + URL + 协议版本号
​			实例： POST + /chapter17/user + HTTP/1.1
​			请求方法：
​				GET：获取数据
​				POST：上传数据（表单格式，json格式）
​				PUT：修改数据
​				DELETE： 用于删除数据

#### 	2、请求头

​			格式： key ：value
​			可以有很多个键值对（包含协议自带，也包含用户自定义的）
​			常见重要头：
​				Accept : 接收数据的格式
​				User-Agent : 描述用户浏览器的信息
​				Connection： Keep-Alive (长链接)， Close（短连接）
​				Aceept-Encoding : gzip, xxx , 描述可以接受的编码
​				Cookie: 由服务器设置的key=value数据，客户端下次请求的时候可以							携带过来
​				Content-Type:
​					appliction/-form(表示上传的数据是表单格式),
​					application/json(表示body的数据是json格式)
​			用户可以自定义的：
​				name : Duke
​				age : 18

#### 	3、空行

​			告诉服务器，请求头结束了，用于分割

#### 	4、请求包体(可选的)

​			一般在POST方法时，会配套提供BODY
​			在GET的时候也可以提供BODY，但是这样容易让人混淆，建议不要这			样使用
​			上传两种数据格式：
​				表单： 姓名，性别，年龄
​				json数据格式



### 6、HTTP响应消息格式：

#### 第一部分：状态行

​			协议格式： 协议版本号 + 状态码 + 状态描述

​			实例1：HTTP/1.1 + 200 + OK

​			实例2：HTTP/1.1 +404 + not found

​			常用的状态码：

				1xx ===》客户端可以即系发送请求（一般感知不到）
				2xx ===》正常访问， 200
				3xx ===》重定向
				4xx
					401 ===》 未授权 not authorized
					404 ===> Not found
				5xx
					501 ==> Internal Error （服务器内部错误）
#### 第二部分：响应头

​			Content-Type : application/json
​			Server: Apache
​			Data : Mon, 12 Sep …

#### 第三部分：空行

​			用于分割，表示下面没有响应头了

#### 第四部分：响应包体

​			通常是返回json数据

