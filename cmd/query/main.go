// Package main API entrypoint.
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/internal/infra/db"
	"github.com/julioc98/delivery/pkg/cache"
	"github.com/julioc98/delivery/pkg/database"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/julioc98/delivery/pkg/pb"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	// Create connections.
	conn, err := database.Conn()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := cache.RedisConn()
	if err != nil {
		log.Fatal(err)
	}

	natsClient, err := event.NatsConn()
	if err != nil {
		log.Fatal(err)
	}

	// Create NATS producer.
	natsProducer := event.NewNatsProducer(natsClient)

	// Create dependencies.
	postGisRepo := db.NewDeliveryPostGISRepository(conn)

	deliveryRepository := db.NewRedisRepositoryDecorator(redisClient, postGisRepo)

	deliveryUseCase := app.NewDeliveryUseCase(deliveryRepository, natsProducer)

	// Create handlers.
	grpcHandler := api.NewGrpcHandler(deliveryUseCase)

	// Start server.
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterQueryServiceServer(s, grpcHandler)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
