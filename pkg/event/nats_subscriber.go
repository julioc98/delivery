// Package event provides the event service.
package event

import "github.com/nats-io/nats.go"

// NatsSubscriber is a producer that uses NATS as a message broker.
type NatsSubscriber struct {
	nc *nats.Conn
}

// NewNatsSubscriber creates a new NatsSubscriber.
func NewNatsSubscriber(nc *nats.Conn) *NatsSubscriber {
	return &NatsSubscriber{nc: nc}
}

// Subscribe subscribes to a NATS subject.
func (p *NatsSubscriber) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return p.nc.Subscribe(subject, cb)
}
