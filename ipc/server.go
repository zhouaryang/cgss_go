package ipc

import (
	"encoding/json"
	"fmt"
)

//简单的IPC框架

type Request struct { //请求内容结构体
	Method string "method"
	Params string "params"
}

type Response struct { //响应内容结构体
	Code string "code"
	Body string "body"
}

type Server interface { //确定了实现服务器的统一接口 、、服务端接口
	Name () string //名称
	Handle(method,params string) *Response //处理方法参数
}

type IpcServer struct {//IpcServer服务内容结构体
	Server //服务端接口，可以定义多个服务层
}

func NewIpcServer (server Server) *IpcServer { //创建IpcServer方法，返回值是*ipcserver
	return &IpcServer{server} //使用花括号可以给类型中的变量赋值 相当于给IpServer结构体赋值,Server:server

}
//结构体 server 有一个 方法 connect ， 这个方法的返回值 有两个， 一个 是chan 类型， 一个是string 类型
func (server *IpcServer)Connect () chan string {
	session := make(chan string ,0) //创建一个session是string类型的通道，长度为0的缓冲区

	go func(c chan string) {  // 另起轻量级线程实现goroutine，匿名函数，传入参数c，c为string类型的chan通道
		for {
			request := <-c //从c中取出数据到request

			if request == "CLOSE" { //关闭该连接 判断从通道取到的是否是close
				break
			}
			var req Request //定义req为Request类型
			err := json.Unmarshal([]byte(request),&req)
			if err != nil {
				fmt.Println("Invalid request format:",request)
			}

			resp := server.Handle(req.Method,req.Params)

			b,err := json.Marshal(resp)

			c <- string(b) //返回结果
		}
	}(session)

	fmt.Println("A new session has been created successfully.")

	return session
}

