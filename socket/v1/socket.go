package main

import (
	"fmt"
	"net"
)

var maps map[string]net.Conn

func main() {
	// map集合用于存储连接
	maps = make(map[string]net.Conn)
	// 监听8080端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("连接出错：%v\n", err)
	}

	// 持续接收请求
	for {
		c, _ := listen.Accept()
		// 存储连接的地址和net.conn对象
		maps[c.RemoteAddr().String()] = c
		// 开启一个协程处理发送的信息
		go readAndSend(c)
	}

}

// 读和发送
func readAndSend(c net.Conn) {
	c.Write([]byte("连接成功"))
	for {
		var buff = make([]byte, 1000)
		_, _ = c.Read(buff)
		fmt.Printf("接收到来自"+c.RemoteAddr().String()+"的数据======>:", string(buff))

		// 给所有用户广播发送
		for str := range maps {
			if str != c.RemoteAddr().String() {
				writeToBroadcast(maps[str], buff, c.RemoteAddr().String())
			}
		}
	}
}

// 广播发送
func writeToBroadcast(c net.Conn, buff []byte, remoteAddr string) {
	str := remoteAddr + ":" + string(buff)
	c.Write([]byte(str))
}
