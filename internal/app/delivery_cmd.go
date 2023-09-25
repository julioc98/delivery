// Package mdg represents the message driven gateway.
package app

import (
	"fmt"
)

// CreateDriverPositionSubject represents a create driver position subject.
const CreateDriverPositionSubject = "create.driver.position"

// DeliveryCmd represents a NATS message producer.
type DeliveryCmd struct {
	publisher Publisher
}

// NewDeliveryCmd creates a new DeliveryCmd.
func NewDeliveryCmd(publisher Publisher) *DeliveryCmd {
	return &DeliveryCmd{publisher: publisher}
}

// SaveDriverPosition saves a driver position.
func (p *DeliveryCmd) SaveDriverPosition(driverID uint64, latitude, longitude float64) error {
	err := p.publisher.Publish(CreateDriverPositionSubject, []byte(
		fmt.Sprintf(
			`{ "driverId":%d, "latitude":%f,"longitude":%f }`,
			driverID, latitude, longitude)))
	if err != nil {
		return err
	}

	return nil
}
