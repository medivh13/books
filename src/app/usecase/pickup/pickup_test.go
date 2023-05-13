package pickup

import (
	mockDTO "books/mocks/app/dto/pickup"
	mockPublisher "books/mocks/infra/publisher"
	dto "books/src/app/dto/pickup"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockPickUpUseCase struct {
	mock.Mock
}

type PickUpUseCaseList struct {
	suite.Suite

	mockDTO       *mockDTO.MockPickupDTO
	mockPublisher *mockPublisher.MockPublisher
	useCase       PickUpUCInterface

	dtoCreate *dto.ReqPickupDTO
}

func (suite *PickUpUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockPickupDTO)
	suite.mockPublisher = new(mockPublisher.MockPublisher)
	suite.useCase = NewPickUpUseCase(suite.mockPublisher)

	suite.dtoCreate = &dto.ReqPickupDTO{
		Date: "12-05-2023",
		User: "jody",
		Information: []*dto.Information{
			&dto.Information{
				Title: "test",
			},
		},
	}

}

func (u *PickUpUseCaseList) TestCreateSuccess() {
	newData, _ := json.Marshal(u.dtoCreate)
	u.mockPublisher.Mock.On("Nats", newData, "books").Return(nil)

	err := u.useCase.Create(u.dtoCreate)
	u.Equal(nil, err)
}

func (u *PickUpUseCaseList) TestCreateFail() {
	newData, _ := json.Marshal(u.dtoCreate)
	u.mockPublisher.Mock.On("Nats", newData, "books").Return(errors.New(mock.Anything))

	err := u.useCase.Create(u.dtoCreate)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *PickUpUseCaseList) TestAddDataPickUp() {

	u.useCase.AddDataPickUp(u.dtoCreate)

}

func (u *PickUpUseCaseList) TestGetDataPickUp() {

	u.useCase.GetDataPickUp()

}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(PickUpUseCaseList))
}
