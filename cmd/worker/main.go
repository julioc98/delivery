// Package main API entrypoint.
package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/db"
	"github.com/julioc98/delivery/internal/infra/msg"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func main() {
	// Create connections.
	conn, err := dbConn()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := redisConn()
	if err != nil {
		log.Fatal(err)
	}

	natsClient, err := natsConn()
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

func dbConn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func redisConn() (*redis.Client, error) {
	opt, err := redis.ParseURL("redis://@localhost:6379/0")
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	return client, nil
}

func natsConn() (*nats.Conn, error) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
