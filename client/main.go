package main

import (
	"log"

	pb "github.com/karokojnr/grpc-golang/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id) // valid
	// readBlog(c, "aNonExistingId")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
