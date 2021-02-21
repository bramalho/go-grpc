package main

import (
	"fmt"
	"github.com/bramalho/go-grpc/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{Body: "Hello from the Client"}
	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
