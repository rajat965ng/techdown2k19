package main

import (
	"context"
	"fmt"
	"github.com/rajat965ng/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("This is a client ...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("The client unable to connect server: %v",err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("The client established connection: %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient)  {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rajat",
			LastName: "Nigam",
		},
	}

	res,err := c.Greet(context.Background(),req)
	if err!=nil {
		log.Fatalf("The error caused while calling greet is %v",err)
	}
	log.Printf("The result of greet call is : %v", res.GetResult())

}