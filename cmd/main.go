package main

import (
	"fmt"

	"github.com/golangproject/delivery/grpc"

	"github.com/golangproject/repository/memcache"
)

func main() {
	fmt.Println("Server Starting")
	memcache := memcache.New()
	grpc.New(memcache)
}
