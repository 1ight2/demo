package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"time"
)

func main() {
	go func() {
		ln, err := net.Listen("tcp", "127.0.0.1:9999")
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			//使用 Accept() 方法等待并接受一个新的连接
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}

			go func() {
				//处理conn
				b := make([]byte, 1234)
				_, err := conn.Read(b)
				if err != nil {
					log.Println(err)
					return
				}
				//log.Fatalln(string(b[:n]))
			}()
		}
	}()

	time.Sleep(100 * time.Millisecond)

	cc, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() //资源释放

	//分片
	for i := 0; i < 1024; i++ {
		time.Sleep(100 * time.Millisecond)
		cc.Write([]byte(strings.Repeat("s", 1024*1024)))
		select {
		case <-ctx.Done():
			fmt.Println("closed", ctx.Err())
			return
		default:

		}
	}
	time.Sleep(100 * time.Millisecond)

}
