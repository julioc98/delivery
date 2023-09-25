// Package event provides the event service.
package event

import "github.com/nats-io/nats.go"

// NatsProducer is a producer that uses NATS as a message broker.
type NatsProducer struct {
	nc *nats.Conn
}

// NewNatsProducer creates a new NatsProducer.
func NewNatsProducer(nc *nats.Conn) *NatsProducer {
	return &NatsProducer{nc: nc}
}

// Publish publishes a message to a NATS subject.
func (p *NatsProducer) Publish(subject string, data []byte) error {
	return p.nc.Publish(subject, data)
}
