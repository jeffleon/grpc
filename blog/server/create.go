package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jeffleon/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create blog was invoked with %v\n", in)
	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Title,
	}
	fmt.Println("OK ", data)
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println("Errr", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Canot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
