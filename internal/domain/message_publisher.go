package domain

type MessagePublisher interface {
	Publish(body []byte) error
}
