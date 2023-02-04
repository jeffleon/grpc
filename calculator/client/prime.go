package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		PrimeRequest: 720,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while readinf the stream: %v\n", err)
		}

		log.Printf("Number: %d\n", msg.PrimeNumber)

	}
}
