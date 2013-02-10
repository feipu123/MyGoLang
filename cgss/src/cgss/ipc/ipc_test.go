package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"ECHO:" + method + params, ""}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("From Client1", "xx")
	resp2, _ := client2.Call("From Client2", "xx")

	if resp1.Code != "ECHO:From Client1xx" || resp2.Code != "ECHO:From Client2xx" {
		t.Error("IpcClient.Call failed. resp1:", resp1.Code, "resp2:", resp2.Code)
	}

	client1.Close()
	client2.Close()
}
