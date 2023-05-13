package email

import (
	dto "books/src/app/dto/pickup"

	"github.com/stretchr/testify/mock"
)

type MockPickupDTO struct {
	mock.Mock
}

func NewMockBooksDTO() *MockPickupDTO {
	return &MockPickupDTO{}
}

var _ dto.PickUpDTOInterface = &MockPickupDTO{}

func (m *MockPickupDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
