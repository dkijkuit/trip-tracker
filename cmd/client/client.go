package main

import (
	"context"
	"log"
	triptrackerserver "trip-tracker/grpc/triptracker/proto"

	gRPC "google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:5000"
	conn, e := gRPC.Dial(serverAddress, gRPC.WithInsecure())

	if e != nil {
		panic(e)
	}

	client := triptrackerserver.NewTripServiceClient(conn)

	trip := triptrackerserver.Trip{
		Id:                 1,
		StartAddress:       "Rotterdam",
		DestinationAddress: "Utrecht",
		Distance:           40,
	}

	req := triptrackerserver.AddTripRequest{
		Trip: &trip,
	}

	response, err := client.AddTrip(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	log.Println(response)
}
