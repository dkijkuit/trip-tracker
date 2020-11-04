package triptrackerservice

import (
	"context"
	"log"

	pb "trip-tracker/grpc/triptracker/proto"
)

var trips []*pb.Trip

// TripServiceServerImpl implements the interface from triptracker.pb.go
type TripServiceServerImpl struct {
}

// NewTripServiceServerImpl returns a pointer to the implementation
func NewTripServiceServerImpl() *TripServiceServerImpl {
	return &TripServiceServerImpl{}
}

//AddTrip service method for adding a new
func (*TripServiceServerImpl) AddTrip(ctx context.Context, req *pb.AddTripRequest) (*pb.AddTripResponse, error) {
	trips = append(trips, req.Trip)
	response := pb.AddTripResponse{
		Added: true,
	}

	log.Println("Added trip, peopleCount:", req.Trip.PeopleCount)

	return &response, nil
}

//GetTrips retrieves all trips from the slice
func (*TripServiceServerImpl) GetTrips(ctx context.Context, req *pb.GetTripsRequest) (*pb.GetTripsResponse, error) {
	return &pb.GetTripsResponse{
		Trips: trips,
	}, nil
}
