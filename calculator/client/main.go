package main

import (
	"log"

	pb "github.com/jeffleon/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrime(c)
	//doAVG(c)
	doSqrt(c, -2)

}
