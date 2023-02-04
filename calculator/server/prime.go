package main

import (
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime function was invoked with %v\n", in)
	num := in.PrimeRequest
	var i uint32
	for i = 2; num > 1; {
		if num%i == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeNumber: i,
			})
			num = num / i
		} else {
			i++
		}
	}

	return nil
}
