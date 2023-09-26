package app

import (
	"encoding/json"

	"github.com/julioc98/delivery/internal/domain"
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
	driverPosition := domain.DriverPosition{
		DriverID: driverID,
		Location: domain.Point{Lat: latitude, Long: longitude},
	}

	msg, err := json.Marshal(driverPosition)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(CreateDriverPositionSubject, msg)
	if err != nil {
		return err
	}

	return nil
}
