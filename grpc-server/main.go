package main

import (
	"fmt"
	"log"
	api "massive-search/grpc-api"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.NewServer()

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the StoredWords service to the server
	api.RegisterStoredWordsServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
