package event

import "github.com/nats-io/nats.go"

// NatsConn creates a NATS connection.
func NatsConn() (*nats.Conn, error) {
	// Connect to a server
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		return nil, err
	}

	return nc, nil
}
