package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	defer conn.Close()

	// 读取用户输入的名称
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入你的名称: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// 发送名称到服务器
	_, err = conn.Write([]byte(name + "\n"))
	if err != nil {
		fmt.Println("发送名称时出错:", err)
		return
	}

	// 等待另一个玩家连接
	fmt.Println("等待另一个玩家连接...")
	_, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取服务器响应时出错:", err)
		return
	}

	// 开始游戏循环
	for {
		fmt.Println("请输入你的选择 (石头/剪刀/布): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// 发送游戏选择到服务器
		_, err = conn.Write([]byte(choice + "\n"))
		if err != nil {
			fmt.Println("发送选择时出错:", err)
			return
		}

		// 读取服务器的响应（游戏结果）
		result, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("读取游戏结果时出错:", err)
			return
		}
		fmt.Print("游戏结果: ", result)
	}
}
