package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {}

func (*Server) Sum(ctx context.Context, req *calcpb.Request) (*calcpb.Response, error){
	res := &calcpb.Response{
		Result: req.GetInput().GetFirstNum() + req.GetInput().GetSecondNum(),
	}
	return res, nil
}

func main() {
	fmt.Println("Starting Server ....")
	lis,err := net.Listen("tcp","0.0.0.0:50051")
	if err!=nil {
		log.Fatalf("Unable to start server: %v",err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s,&Server{})
	if err = s.Serve(lis); err!=nil {
		log.Fatalf("Server is unable to serve: %v",err)
	} 
}
