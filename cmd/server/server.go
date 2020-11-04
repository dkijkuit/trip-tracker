package main

import (
	"fmt"
	"log"
	"net"

	triptrackerserver "trip-tracker/grpc/triptracker/proto"
	triptrackerservice "trip-tracker/service/triptracker"

	"google.golang.org/grpc"
)

func main() {
	startGrpcTripServer()
}

func startGrpcTripServer() {
	log.Println("Starting gRPC server instance....")
	tripserver := triptrackerservice.NewTripServiceServerImpl()
	lis := getNetListener(5000)

	grpcServer := grpc.NewServer()
	triptrackerserver.RegisterTripServiceServer(grpcServer, tripserver)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
