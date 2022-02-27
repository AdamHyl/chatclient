package main

import (
	"flag"
	"fmt"

	"github.com/AdamHyl/chatclient/client"
)

//尝试从终端命令行解析IP和Port创建客户端
var serverIp string
var serverPort int

//文件的初始化函数
//命令的格式  ./client.exe -ip 127.0.0.1 -port 8888
func init() {
	//属于初始化工作，一般放在init中
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 3063, "set server port(default:8888)")
}

func main() {
	//通过命令行解析
	flag.Parse()
	c := client.NewClient(serverIp, serverPort)
	//c := newClient("127.0.0.1", 8888)
	if c == nil {
		fmt.Println("------- connect server error------")
		return
	}

	fmt.Println("-------- connect server success ------")

	fmt.Println("ready to process transactions......")
	c.Run()

}
