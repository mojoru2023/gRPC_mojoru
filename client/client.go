package main

import (
	"context"
	"flag"
	pb "gRPC_mojoru/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 默认数据 也支持在控制台自定义
const (
	defaultName = "world"
)

// 监听地址和传入的name
var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	// name = flag.String("name", defaultName, "Name to greet")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	//通过gRPC.Dial()方法建立服务连接
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//连接要记得关闭
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	//实例化客户端连接
	c := pb.NewHelloClient(conn)

	//设置请求上下文，因为是网络请求，我们需要设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//客户端调用在proto中定义的SayHello()rpc方法，发起请求，接收服务端响应
	r, err := c.Say(ctx, &pb.SayRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
