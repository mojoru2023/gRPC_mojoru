package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC_mojoru/proto" //导入我们在protos文件中定义的服务

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 定义一个结构体，作用是实现mojoru中的HelloServer
type server struct {
	proto.UnimplementedHelloServer
}

func (s *server) Say(ctx context.Context, req *proto.SayRequest) (*proto.SayResponse, error) {
	fmt.Println("request:", req.Name)
	return &proto.SayResponse{Message: "Hello " + req.Name}, nil
}

// 定义端口号 支持启动的时候输入端口号
var (
	port = flag.Int("port", 50051, "The server port")
)
//客户端和服务端代码中的flag.Parse的作用是：
// 支持我们在终端控制台自定义输入参数，如果没有输入的话，
// 使用程序中设置的默认参数，比如客户端的name，在代码中是这么定义的：
// name = flag.String("name", "world", "Name to greet")
// 我们在终端输入如下命令：

// go run main.go --name 王中阳

func main() {
	//解析输入的端口号 默认50051
	flag.Parse()
	//tcp协议监听指定端口号
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	//实例化gRPC服务
	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})
	reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 50051...")
	// 启动服务
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
