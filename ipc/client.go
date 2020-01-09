package ipc

import (
	"encoding/json"
)

type IPCClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IPCClient {
	c := server.Connect()
	return &IPCClient{c}
}
//函数call，是客户端ipcclient定义的client的方法，传入参数为string类型的method和params，返回参数为 *Response类型的resp和error
func (client *IPCClient) Call (method , params string) (resp *Response,err error){
	//获取request请求结构体，参数为method，params
	req := &Request{method,params}
	//定义type类型数组b
	var b []byte
	b, err = json.Marshal(req)//将req序列化为json格式串

	if err != nil {
		return
	}
	client.conn <- string(b)//将json串b传入通道client.conn
	str := <- client.conn // 等待返回值//获取client.conn通道字符串

	var resp1 Response//定义response类型结构体resp1
	err = json.Unmarshal([]byte(str),&resp1)//反序列化str为结构体resp1
	resp = &resp1//将resp1赋值给resp
	return
}

func (client *IPCClient)Close(){ // 添加成员函数//ipc客户端client定义一个close函数
	client.conn <- "CLOSE"//将close字符串扔进client.conn通道
}