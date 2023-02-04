package main

import (
	"context"
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: number,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.InvalidArgument {
				log.Println("we probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
