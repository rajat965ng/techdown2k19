package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Starting client ...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while dialing to server: %v", err)
	}
	defer conn.Close()

	cc := calcpb.NewCalculatorServiceClient(conn)
	doUnary(cc)
	doServerStream(cc)
	doClientStreaming(cc)
}

func doUnary(cc calcpb.CalculatorServiceClient) {
	req := &calcpb.Request{
		Input: &calcpb.Input{
			FirstNum:  7,
			SecondNum: 13,
		},
	}

	resp, err := cc.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in computing sum: %v", err)
	}
	log.Printf("The sum is: %v", resp.GetResult())
}

func doServerStream(cc calcpb.CalculatorServiceClient) {
	req := &calcpb.MeteoricNumber{
		Value: 120,
	}
	resp, err := cc.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in computing sum: %v", err)
	}
	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("The client recieve error from prime number stream: %v", err)
		}
		log.Printf("The result is: %v", msg.GetValue())
	}
}

func doClientStreaming(cc calcpb.CalculatorServiceClient) {
	arr := [4]calcpb.NumSeries{
		calcpb.NumSeries{
			Num: 2,
		},
		calcpb.NumSeries{
			Num: 3,
		},
		calcpb.NumSeries{
			Num: 4,
		},
		calcpb.NumSeries{
			Num: 5,
		},
	}

	stream, _ := cc.SumSeries(context.Background())
	for _, v := range arr {
		err := stream.Send(&v)
		if err != nil {
			log.Fatalf("Error while computing sum of series: %v", err)
		}
	}
	sum, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing stream: %v", err)
	}
	log.Printf("The sum of num series is %v", sum.Num)

}