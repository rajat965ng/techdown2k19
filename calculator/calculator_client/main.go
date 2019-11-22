package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Starting client ...")

	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile,"")
	if sslErr!=nil {
		log.Fatalf("Unable to load certs %v", sslErr)
	}

	opts := grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Error while dialing to server: %v", err)
		return
	}
	defer conn.Close()

	cc := calcpb.NewCalculatorServiceClient(conn)
	doUnary(cc)
	doServerStream(cc)
	doClientStreaming(cc)
	doBiDirectionalStreaming(cc)
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
		Value: 12,
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

func doBiDirectionalStreaming(cc calcpb.CalculatorServiceClient) {
	arr := [3]calcpb.Account{
		{Principal: 2000,
			Rate:        8,
			TimeInYears: 3},
		{Principal: 3000,
			Rate:        7,
			TimeInYears: 4},
		{Principal: 5000,
			Rate:        6,
			TimeInYears: 5},
	}
	stream, _ := cc.CalculateInterest(context.Background())

	buff := make(chan struct{})
	go func(chan struct{}) {
		for _, v := range arr {
			log.Printf("Sending the Account record: %v",&v)
			err := stream.Send(&v)
			time.Sleep(1*time.Second)
			if err != nil {
				log.Printf("Error while sending account over stream: %v", err)
			}
		}
		stream.CloseSend()
	}(buff)

	go func(chan struct{}) {
		for {
			interest, err := stream.Recv()
			if err == io.EOF {
				close(buff)
				break
			}
			if err != nil {
				log.Printf("Error while recieving simple interest: %v", err)
			}
			log.Printf("The simple interest is: %v", interest.GetInterest())
		}
	}(buff)

	<-buff

}
