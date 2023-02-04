package main

import (
	"context"
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	sum, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  8,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("could not sum: %v\n", err)
	}

	log.Printf("Sum %d\n", sum.Result)

}
