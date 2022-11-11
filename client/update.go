package main

import (
	"context"
	"log"

	pb "github.com/karokojnr/grpc-golang/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog function invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Kennedy",
		Title:    "gRPC - Golang",
		Content:  "https://grpc.io/docs/languages/go/basics/",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened when updating: %v\n", err)
	}
	log.Println("Blog was updated!")
}
