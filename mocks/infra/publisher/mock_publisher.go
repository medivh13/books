package mock_worker

import (
	publisher "books/src/infra/broker/nats/publisher"

	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

func NewMockWorker() *MockPublisher {
	return &MockPublisher{}
}

var _ publisher.PublisherInterface = &MockPublisher{}

func (m *MockPublisher) Nats(data []byte, subject string) error {
	args := m.Called(data, subject)
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}

func (m *MockPublisher) InitNats() {
}
