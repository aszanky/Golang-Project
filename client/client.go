package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golangproject/delivery/contentpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Content Client Testing")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := contentpb.NewContentServiceClient(cc)

	// create Content
	fmt.Println("Creating the Content")
	var dummydata []*contentpb.Content
	for i := 1; i < 6; i++ {
		dataContent := &contentpb.Content{
			Id:      int32(i),
			Content: fmt.Sprintf("Content dummy %d", i),
		}
		dummydata = append(dummydata, dataContent)
	}

	for _, v := range dummydata {
		createContentRes, err := c.CreateContent(context.Background(), &contentpb.CreateContentReq{Content: v})
		if err != nil {
			log.Fatalf("Unexpected error: %v", err)
		}
		fmt.Printf("Content has been created: %v\n", createContentRes)
	}

	//GetALL data
	getAllData, err := c.GetAllContent(context.Background(), &contentpb.GetAllContentReq{})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Getting All data: %v", getAllData)

	//Delete Data
	deleteSuccess, err := c.DeleteContent(context.Background(), &contentpb.DeleteContentReq{Content: &contentpb.Content{
		Id: 1,
	}})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Delete Data %v", deleteSuccess)
}
