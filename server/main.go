package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/asyrawi/proto_example/model"
	"google.golang.org/grpc"
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

	log.Printf("Server Running on :%v", list.Addr().String())

	log.Fatal(grpc_server.Serve(list))

}

func (s *Server) GetUser(ctx context.Context, userId *pb.UserId) (*pb.User, error) {
	fmt.Println(userId)
	return &pb.User{
		Name:        "Hanan",
		Age:         12,
		Address:     "Tomoni",
		PhoneNumber: "0851234123123",
	}, nil
}
func (s *Server) GetAllUser(context.Context, *pb.Void) (*pb.Users, error) {
	// You Can Use Database Call In Here
	users := &pb.Users{
		List: []*pb.User{
			{
				Name:        "Hanan",
				Age:         12,
				Address:     "Tomoni",
				PhoneNumber: "0851234123123",
			},
			{
				Name:        "Asyrawi",
				Age:         12,
				Address:     "Tomoni",
				PhoneNumber: "0851234123123",
			},
		},
	}

	return users, nil
}
