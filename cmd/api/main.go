// Package main API entrypoint.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/internal/infra/db"
	"github.com/julioc98/delivery/pkg/cache"
	"github.com/julioc98/delivery/pkg/database"
	"github.com/julioc98/delivery/pkg/event"

	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

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
