package main

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/rajat965ng/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"net"
)

type Server struct {}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("The greet function was invoked with req: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello "+firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res,nil
}

func main() {
	fmt.Println("Hello Server....")
	lis,err := net.Listen("tcp","0.0.0.0:50051")
	if err!=nil{
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s,&Server{})
	if err=s.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
