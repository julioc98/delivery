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
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	// Create dependencies.
	conn, err := dbConn()
	if err != nil {
		log.Fatal(err)
	}

	postGisRepo := db.NewDeliveryPostGISRepository(conn)

	redisClient, err := redisConn()
	if err != nil {
		log.Fatal(err)
	}

	deliveryRepository := db.NewRedisRepositoryDecorator(redisClient, postGisRepo)

	deliveryUseCase := app.NewDeliveryUseCase(deliveryRepository)

	// Create handlers.
	deliveryRestHandler := api.NewDeliveryRestHandler(r, deliveryUseCase)
	deliveryRestHandler.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Println("Starting server on port 3000...")

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		_ = conn.Close()

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
