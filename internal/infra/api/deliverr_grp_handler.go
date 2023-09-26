package api

import (
	"context"

	"github.com/julioc98/delivery/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// QueryServiceServer represents a query service server.
type QueryServiceServer struct {
	pb.UnimplementedQueryServiceServer
	qry DeliveryQry
}

// NewGrpcHandler creates a new gRPC handler.
func NewGrpcHandler(qry DeliveryQry) *QueryServiceServer {
	return &QueryServiceServer{qry: qry}
}

// FindDriverPosition finds a driver position.
func (s *QueryServiceServer) FindDriverPosition(
	_ context.Context, request *pb.FindDriverPositionRequest,
) (
	*pb.FindDriverPositionResponse, error,
) {
	position, err := s.qry.FindDriverPosition(request.DriverId)
	if err != nil {
		return nil, err
	}

	return &pb.FindDriverPositionResponse{
		Position: &pb.DriverPosition{
			DriverId:  request.DriverId,
			Latitude:  position.Location.Lat,
			Longitude: position.Location.Long,
			Timestamp: timestamppb.New(position.Timestamp),
		},
	}, nil
}

// HistoryDriverPosition finds a driver position history.
func (s *QueryServiceServer) HistoryDriverPosition(
	_ context.Context, request *pb.HistoryDriverPositionRequest,
) (
	*pb.HistoryDriverPositionResponse, error,
) {
	positions, err := s.qry.HistoryDriverPosition(request.DriverId)
	if err != nil {
		return nil, err
	}

	response := &pb.HistoryDriverPositionResponse{}
	for _, position := range positions {
		response.Positions = append(response.Positions, &pb.DriverPosition{
			DriverId:  request.DriverId,
			Latitude:  position.Location.Lat,
			Longitude: position.Location.Long,
			Timestamp: timestamppb.New(position.Timestamp),
		})
	}

	return response, nil
}

// FindDriversNearby finds drivers nearby.
func (s *QueryServiceServer) FindDriversNearby(
	_ context.Context, request *pb.FindDriversNearbyRequest,
) (
	*pb.FindDriversNearbyResponse, error,
) {
	positions, err := s.qry.GetDriversNearby(
		request.Latitude, request.Longitude, int(request.Radius),
	)
	if err != nil {
		return nil, err
	}

	response := &pb.FindDriversNearbyResponse{}
	for _, position := range positions {
		response.Positions = append(response.Positions, &pb.DriverPosition{
			DriverId:  position.DriverID,
			Latitude:  position.Location.Lat,
			Longitude: position.Location.Long,
			Timestamp: timestamppb.New(position.Timestamp),
		})
	}

	return response, nil
}
