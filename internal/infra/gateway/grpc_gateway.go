// Package gateway represents the gateway layer.
package gateway

import (
	"context"

	"github.com/julioc98/delivery/internal/domain"
	"github.com/julioc98/delivery/pkg/pb"
)

// GrpcGateway represents a gRPC gateway.
type GrpcGateway struct {
	cli pb.QueryServiceClient
}

// NewGrpcGateway creates a new gRPC gateway.
func NewGrpcGateway(cli pb.QueryServiceClient) *GrpcGateway {
	return &GrpcGateway{cli: cli}
}

// FindDriverPosition finds a driver position.
func (g *GrpcGateway) FindDriverPosition(driverID uint64) (domain.DriverPosition, error) {
	req := &pb.FindDriverPositionRequest{DriverId: driverID}
	ctx := context.Background()

	res, err := g.cli.FindDriverPosition(ctx, req)
	if err != nil {
		return domain.DriverPosition{}, err
	}

	dp := domain.DriverPosition{
		DriverID: driverID,
		Location: domain.Point{
			Lat:  res.Position.Latitude,
			Long: res.Position.Longitude,
		},
		Timestamp: res.Position.Timestamp.AsTime(),
	}

	return dp, nil
}

// HistoryDriverPosition finds a driver position history.
func (g *GrpcGateway) HistoryDriverPosition(driverID uint64) ([]domain.DriverPosition, error) {
	req := &pb.HistoryDriverPositionRequest{DriverId: driverID}
	ctx := context.Background()

	res, err := g.cli.HistoryDriverPosition(ctx, req)
	if err != nil {
		return nil, err
	}

	dps := []domain.DriverPosition{}

	for _, p := range res.Positions {
		dp := domain.DriverPosition{
			DriverID: driverID,
			Location: domain.Point{
				Lat:  p.Latitude,
				Long: p.Longitude,
			},
			Timestamp: p.Timestamp.AsTime(),
		}

		dps = append(dps, dp)
	}

	return dps, nil
}

// GetDriversNearby finds drivers nearby.
func (g *GrpcGateway) GetDriversNearby(latitude, longitude float64, radius int) ([]domain.DriverPosition, error) {
	req := &pb.FindDriversNearbyRequest{
		Latitude:  latitude,
		Longitude: longitude,
		Radius:    int32(radius),
	}
	ctx := context.Background()

	res, err := g.cli.FindDriversNearby(ctx, req)
	if err != nil {
		return nil, err
	}

	dps := []domain.DriverPosition{}

	for _, p := range res.Positions {
		dp := domain.DriverPosition{
			DriverID: p.DriverId,
			Location: domain.Point{
				Lat:  p.Latitude,
				Long: p.Longitude,
			},
			Timestamp: p.Timestamp.AsTime(),
		}

		dps = append(dps, dp)
	}

	return dps, nil
}
