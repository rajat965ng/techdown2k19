package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting client ...")
	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err!=nil {
		log.Fatalf("Error while dialing to server: %v", err)
	}
	defer conn.Close()

	cc := calcpb.NewCalculatorServiceClient(conn)
	doUnary(cc)
}

func doUnary(cc calcpb.CalculatorServiceClient)  {
	req:=&calcpb.Request{
		Input: &calcpb.Input{
			FirstNum:7,
			SecondNum:13,
		},
	}

	resp,err := cc.Sum(context.Background(),req)
	if err!=nil{
		log.Fatalf("Error in computing sum: %v",err)
	}
	log.Printf("The sum is: %v",resp.GetResult())
}