package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func doAVG(c pb.CalculatorServiceClient) {
	reqs := []*pb.AvgRequest{
		{Number: 3},
		{Number: 5},
		{Number: 7},
		{Number: 10},
	}
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while reciving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %f\n", res.Avg)
}
