package main

import (
	"log"

	pb "github.com/jeffleon/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBLog(c, id)
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
