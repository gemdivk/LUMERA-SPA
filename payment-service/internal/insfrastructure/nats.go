package infrastructure

import (
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/domain"
	"github.com/nats-io/nats.go"
)

type NATSAdapter struct {
	Conn *nats.Conn
}

func NewNATSClient(conn *nats.Conn) domain.NATSClient {
	return &NATSAdapter{Conn: conn}
}

func (n *NATSAdapter) Publish(subject string, data []byte) error {
	return n.Conn.Publish(subject, data)
}
