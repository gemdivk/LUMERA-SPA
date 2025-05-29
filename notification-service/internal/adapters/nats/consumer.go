package nats

import (
	"encoding/json"
	"log"

	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain/application"
	"github.com/nats-io/nats.go"
)

type Consumer struct {
	nc      *nats.Conn
	usecase application.EmailUsecase
}

func NewConsumer(nc *nats.Conn, uc application.EmailUsecase) *Consumer {
	return &Consumer{nc: nc, usecase: uc}
}

func (c *Consumer) Subscribe() {
	_, err := c.nc.Subscribe("notifications.email.verification", func(m *nats.Msg) {
		var payload map[string]string
		if err := json.Unmarshal(m.Data, &payload); err != nil {
			log.Println("Failed to parse payload:", err)
			return
		}
		c.usecase.SendVerificationEmail(payload["email"], payload["token"])
	})
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}
}
