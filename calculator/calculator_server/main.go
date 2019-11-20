package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct{}

func (*Server) SumSeries(stream calcpb.CalculatorService_SumSeriesServer) error  {
	sum:=int64(0)
	for  {
		list,err:=stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error is receiving request stream: %v", err)
		}
		sum += list.Num
	}
	stream.SendAndClose(&calcpb.SeriesSum{
		Num: sum,
	})
	return nil
}

func (*Server) PrimeNumberDecomposition(req *calcpb.MeteoricNumber, stream calcpb.CalculatorService_PrimeNumberDecompositionServer) error {

	log.Printf("The number recieved is :%v", req)
	for input, div := req.GetValue(), int64(2); input > 1; {
		if input%div == 0 {
			input = input / div
			stream.Send(&calcpb.PrimeFactor{
				Value: div,
			})
		} else {
			div ++
		}
	}

	return nil
}

func (*Server) Sum(ctx context.Context, req *calcpb.Request) (*calcpb.Response, error) {
	res := &calcpb.Response{
		Result: req.GetInput().GetFirstNum() + req.GetInput().GetSecondNum(),
	}
	return res, nil
}

func main() {
	fmt.Println("Starting Server ....")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Server is unable to serve: %v", err)
	}
}
