package service

type PublisherService interface {
	Push(data []byte) error
}