package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	ADDRESS := "127.0.0.1:5000"
	conn,err := net.Dial("tcp",ADDRESS) // 主动与服务端建立连接
	if err != nil {
		fmt.Printf("dial %s failed; err :%s",ADDRESS,err)
		return
	}
	fmt.Println("客户端开启....")
	var msg string
	var data [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		//发送信息
		fmt.Print("请输入：")
		msg,_ = reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, _ = conn.Write([]byte(msg))

		// 接受信息
		n,err:=conn.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from conn failed, err :",err)
			return
		}
		fmt.Println("收到的回复:",string(data[:n]))
	}
	_ = conn.Close()
}
