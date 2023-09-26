// Package main API entrypoint.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/internal/infra/gateway"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/julioc98/delivery/pkg/pb"
	"github.com/julioc98/delivery/pkg/rpc"

	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	// Create connections.
	natsClient, err := event.NatsConn()
	if err != nil {
		log.Fatal(err)
	}

	gConn, err := rpc.GrpcConn()
	if err != nil {
		log.Fatal(err)
	}

	// Create NATS producer.
	natsProducer := event.NewNatsProducer(natsClient)

	// Create dependencies.

	deliverCmd := app.NewDeliveryCmd(natsProducer)

	c := pb.NewQueryServiceClient(gConn)

	gtGrpc := gateway.NewGrpcGateway(c)

	deliverQry := app.NewDeliveryQry(gtGrpc)

	// Create handlers.
	deliveryRestHandler := api.NewDeliveryRestHandlerByCmdAndQry(r, deliverCmd, deliverQry)
	deliveryRestHandler.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Println("Starting server on port 3001...")

	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
