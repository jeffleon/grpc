package main

import (
	"io"
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	var res uint64 = 0
	var times uint64 = 0

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: float32(res / times),
			})
		}

		if err != nil {
			log.Fatalf("Error while client stream %v\n", err)
		}

		times++
		res += msg.Number
	}
}
