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

func (client *IPCClient) Call (method , params string) (resp *Response,err error){

	req := &Request{method,params}

	var b []byte
	b, err = json.Marshal(req)

	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <- client.conn // 等待返回值

	var resp1 Response
	err = json.Unmarshal([]byte(str),&resp1)
	resp = &resp1

	return
}

func (client *IPCClient)Close(){ // 添加成员函数
	client.conn <- "CLOSE"
}