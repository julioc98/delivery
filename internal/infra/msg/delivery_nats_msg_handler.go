// Package msg implements the message/event driven layer.
package msg

import (
	"encoding/json"
	"log"

	"github.com/julioc98/delivery/internal/app"
	"github.com/julioc98/delivery/internal/domain"
	"github.com/julioc98/delivery/internal/infra/api"
	"github.com/julioc98/delivery/pkg/event"
	"github.com/nats-io/nats.go"
)

// Handler represents a message handler.
type Handler struct {
	sub *event.NatsSubscriber
	us  api.DeliveryUseCase
}

// NewHandler creates a new Handler.
func NewHandler(sub *event.NatsSubscriber, us api.DeliveryUseCase) *Handler {
	return &Handler{sub: sub, us: us}
}

// Register registers a message handler.
func (h *Handler) Register() error {
	_, err := h.sub.Subscribe(app.CreateDriverPositionSubject, func(m *nats.Msg) {
		var v domain.DriverPosition

		unmarshalErr := json.Unmarshal(m.Data, &v)
		if unmarshalErr != nil {
			log.Println("unmarshal err:", unmarshalErr, "data:", string(m.Data))

			return
		}

		saveErr := h.us.SaveDriverPosition(v.DriverID, v.Location.Lat, v.Location.Long)
		if saveErr != nil {
			log.Println("save driver position err:", saveErr, "data:", string(m.Data))

			return
		}
	})

	return err
}
