package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //判断当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999, // 设置flay默认值，否则flag默认为int整型
	}
	//创建链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	//返回客户端
	return client
}

// 监听server回应的消息，直接显示到标准输出
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn) // 永久阻塞监听
}

func (client *Client) PublicChat() {
	//公聊模式

	fmt.Println("please input content, exit for stop")
	var chatMsg = readLine()

	for chatMsg != "exit" {
		//发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		//重新接受下一条消息
		fmt.Println("please input  content, exit for stop")
		chatMsg = readLine()
	}
}

func (client *Client) Run() {
	fmt.Println("please enter your account:")
	account := readLine()
	sendMsg := "1 " + account + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}

	setName := func() {
		name := readLine()
		sendMsg = "2 " + name + "\n"
		_, err = client.conn.Write([]byte(sendMsg))
		if err != nil {
			fmt.Println("conn write err:", err)
			return
		}
	}

	receiveMsg := func() string {
		buf := make([]byte, 4096)
		client.conn.Read(buf)
		s := string(buf)
		// fmt.Println(s)
		return s
	}

	s := receiveMsg()
	set := "please set your name:"
	if strings.HasPrefix(s, set) {
		fmt.Println(set)
		setName()
		s = receiveMsg()
		dif := "please enter a different name:"
		for !strings.HasPrefix(s, "set name ok") {
			fmt.Println(dif)
			setName()
			s = receiveMsg()
		}
		fmt.Println("set name ok")
	}
	go client.DealResponse()

	client.PublicChat()
}

func readLine() string {
	var msg string
	reader := bufio.NewReader(os.Stdin)
	msg, _ = reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	return msg
}
