package domain

type NATSClient interface {
	Publish(subject string, data []byte) error
}
