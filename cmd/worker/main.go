// Package main API entrypoint.
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/db"
	"github.com/julioc98/delivery/internal/infra/msg"
	"github.com/julioc98/delivery/pkg/cache"
	"github.com/julioc98/delivery/pkg/database"
	"github.com/julioc98/delivery/pkg/event"

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

	natsSub := event.NewNatsSubscriber(natsClient)

	// Create dependencies.

	postGisRepo := db.NewDeliveryPostGISRepository(conn)

	deliveryRepository := db.NewRedisRepositoryDecorator(redisClient, postGisRepo)

	deliveryUseCase := app.NewDeliveryUseCase(deliveryRepository, natsProducer)

	// Create handlers.
	natsHandler := msg.NewHandler(natsSub, deliveryUseCase)

	_ = natsHandler.Register()

	// Worker keep alive
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	log.Println("Received interrupt signal. Exiting...")
}
