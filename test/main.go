package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	// 发送请求
	request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}

	// 接收响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)//n是读取到的字节数
	if err != nil {
		fmt.Println("接收响应失败:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("接收到的响应:\n", response)
}