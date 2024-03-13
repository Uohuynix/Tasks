[TOC]



### 1. 创建TCP服务器

#### a. 基本的TCP服务器

```go
gopackage main

import (
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    // 处理连接逻辑
    defer conn.Close()//后续有专门的笔记对该语句进行讲解

    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    fmt.Println("Received message:", string(buffer))
}

func main() {
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on localhost:8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        go handleConnection(conn)
    }
}
```

在这个例子中，我们创建了一个基本的TCP服务器。函数用于在指定的地址和端口上监听TCP连接，然后使用等待客户端连接。一旦有连接建立，会启动一个新的goroutine来处理连接，调用函数。`net.Listen``listener.Accept``handleConnection`

此处建立新的`goroutine`是为了能够响应多个客户端，避免因为单个客户端打开而占用服务器。下文b会有详细解释。

#### b. 处理多个连接

如果你希望服务器能够处理多个连接，可以使用来处理每个连接，以防止一个连接的处理阻塞其他连接的接受。`goroutine`

```go
go// 在主循环中增加goroutine处理连接
for {
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err)
        continue
    }

    go handleConnection(conn)
}
```

### 2. 创建TCP客户端

#### a. 基本的TCP客户端

```go
gopackage main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    message := "Hello, server!"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    fmt.Println("Message sent:", message)
}
```

在这个例子中，我们使用函数连接到指定的TCP服务器地址。然后，我们可以使用连接对象来发送数据。`net.Dial``conn`

### 3. 处理数据流

#### a. 读取和写入数据

TCP连接提供了和方法，用于在连接上进行数据的读取和写入。这些方法接收切片作为参数，允许你读取或写入指定长度的数据。`Read``Write``[]byte`

```go
// 从连接中读取数据
buffer := make([]byte, 1024)
n, err := conn.Read(buffer)
if err != nil {
    fmt.Println("Error reading:", err)
    return
}
fmt.Println("Received data:", string(buffer[:n]))

// 向连接中写入数据
message := "Hello, server!"
_, err := conn.Write([]byte(message))
if err != nil {
    fmt.Println("Error writing:", err)
    return
}
fmt.Println("Message sent:", message)
```

### 4. 完整示例

```go
gopackage main

import (
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    fmt.Println("Received message:", string(buffer[:n]))

    // 向客户端发送响应
    response := "Message received successfully!"
    conn.Write([]byte(response))
}

func main() {
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on localhost:8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        go handleConnection(conn)
    }
}
```

在这个例子中，服务器接收客户端发送的消息并向客户端发送一个简单的响应。