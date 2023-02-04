package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax aws invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while stream %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []uint32{2, 4, 5, 6, 19, 30}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: int32(number),
			})
			time.Sleep(1 * time.Second)

		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving %v\n", err)
				break
			}

			log.Printf("Recieved %+v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
