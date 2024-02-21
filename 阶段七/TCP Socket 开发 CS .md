首先，需要理解以下基础概念：

1. **TCP Socket：** TCP（Transmission Control Protocol）是一种面向连接的协议，它提供可靠的数据传输。Socket 是在网络编程中用于进程间通信的一种机制。
2. **服务器（Server）：** 服务器监听指定的端口，等待客户端连接。一旦连接建立，服务器与客户端之间可以进行双向通信。
3. **客户端（Client）：** 客户端通过指定服务器的地址和端口连接到服务器。一旦连接建立，客户端可以向服务器发送请求，并接收服务器的响应。

以下是一个简单的 Go 语言实现的 C/S 架构示例，其中包含了一个基本的服务器和一个简单的客户端：

**服务器端代码（server.go）：**

```
gopackage main

import (
	"fmt"
	"net"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// 启动一个 goroutine 处理连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 循环接收和处理客户端消息
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		// 处理接收到的消息
		message := string(buffer)
		fmt.Println("Received message:", message)

		// 假设服务器简单地将消息原样返回给客户端
		_, err = conn.Write([]byte("Server: " + message))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			break
		}
	}
}
```

**客户端代码（client.go）：**

```
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type 'exit' to quit.")

	// 循环读取用户输入并发送到服务器
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		// 发送消息到服务器
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			break
		}

		// 接收并打印服务器的响应
		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println("Server response:", string(response))
	}

	fmt.Println("Connection closed.")
}
```