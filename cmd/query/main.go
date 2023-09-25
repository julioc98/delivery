// Package main API entrypoint.
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/internal/infra/db"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

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

	// Create dependencies.

	postGisRepo := db.NewDeliveryPostGISRepository(conn)

	deliveryRepository := db.NewRedisRepositoryDecorator(redisClient, postGisRepo)

	deliveryUseCase := app.NewDeliveryUseCase(deliveryRepository, natsProducer)

	// Create handlers.
	deliveryRestHandler := api.NewDeliveryRestHandler(r, deliveryUseCase)
	deliveryRestHandler.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Println("Starting server on port 3000...")

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		_ = conn.Close()
		_ = redisClient.Close()

		log.Fatal(err)
	}
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
