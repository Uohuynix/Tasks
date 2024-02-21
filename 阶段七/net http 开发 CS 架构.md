**服务器端代码（server.go）：**

```
gopackage main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler) // 设置路由
    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil) // 监听端口
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Client!") // 发送响应内容
}
```

**客户端代码（client.go）：**

```
gopackage main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    response, err := http.Get("http://localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response from server:", string(body))
}
```

在这个示例中，服务器端使用 `http.HandleFunc` 函数将根路由 `/` 映射到 `handler` 函数。当客户端发送请求到服务器的根路径时，服务器会返回 "Hello, Client!" 字符串作为响应。

客户端使用 `http.Get` 函数向服务器发起 GET 请求，并读取响应内容。然后将响应内容打印到控制台上。

要运行这个示例，你需要在命令行中分别执行以下命令：

```
shgo run server.go
shgo run client.go
```

当客户端发送请求后，你将在客户端的控制台上看到服务器发送的响应内容。这个示例虽然简单，但可以帮助你理解基本的 C/S 架构以及使用 Go 的 `net/http` 包进行网络通信的方式。