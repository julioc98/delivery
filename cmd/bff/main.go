// Package main API entrypoint.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/nats-io/nats.go"

	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	// Create connections.
	natsClient, err := natsConn()
	if err != nil {
		log.Fatal(err)
	}

	// Create NATS producer.
	natsProducer := event.NewNatsProducer(natsClient)

	// Create dependencies.

	deliverCmd := app.NewDeliveryCmd(natsProducer)
	deliverQry := app.NewDeliveryQry(natsProducer)

	// Create handlers.
	deliveryRestHandler := api.NewDeliveryRestHandlerByCmdAndQry(r, deliverCmd, deliverQry)
	deliveryRestHandler.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Println("Starting server on port 3000...")

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func natsConn() (*nats.Conn, error) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
