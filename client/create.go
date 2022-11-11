package main

import (
	"context"
	"log"

	pb "github.com/karokojnr/grpc-golang/blog/proto"
)
func createBlog(c pb.BlogServiceClient) string {
	log.Println("-----createBlog was invoked-----")

	blog := &pb.Blog{
		AuthorId: "karokojnr",
		Title: "gRPC Blog",
		Content: "https://grpc.io/",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil{
		log.Fatalf("Unexpected error: %v\n", err)
	}
	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}