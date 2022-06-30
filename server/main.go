package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/asyrawi/proto_example/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserServicesServer
}

var (
	port = flag.Int("port", 2000, "The server port")
)

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return
	}

	// Craete New Server Instances
	grpc_server := grpc.NewServer()

	// Register Service nya di sini
	pb.RegisterUserServicesServer(grpc_server, &Server{})

	log.Printf("Server Running on :%v", list.Addr())

	log.Fatal(grpc_server.Serve(list))

}

func (s *Server) GetUser(context.Context, *pb.UserId) (*pb.Users, error) {
	return nil, nil
}

func (s *Server) GetAllUser(context.Context, *emptypb.Empty) (*pb.Users, error) {
	return nil, nil
}
