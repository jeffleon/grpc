package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jeffleon/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)
	}
}
