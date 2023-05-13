package pickup

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : books
 */

import (
	dto "books/src/app/dto/pickup"
	natsPublisher "books/src/infra/broker/nats/publisher"
	Const "books/src/infra/constants"
	"encoding/json"
	"log"
)

var dataPickup []*dto.ReqPickupDTO

type PickUpUCInterface interface {
	Create(req *dto.ReqPickupDTO) error
	AddDataPickUp(data *dto.ReqPickupDTO)
	GetDataPickUp() []*dto.ReqPickupDTO
}

type pickUpUseCase struct {
	Publisher natsPublisher.PublisherInterface
}

func NewPickUpUseCase(publiser natsPublisher.PublisherInterface) *pickUpUseCase {
	return &pickUpUseCase{
		Publisher: publiser,
	}
}

func (uc *pickUpUseCase) Create(req *dto.ReqPickupDTO) error {
	newData, _ := json.Marshal(req)
	err := uc.Publisher.Nats(newData, Const.BOOKS)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (uc *pickUpUseCase) AddDataPickUp(data *dto.ReqPickupDTO) {

	dataPickup = append(dataPickup, data)
}

func (uc *pickUpUseCase) GetDataPickUp() []*dto.ReqPickupDTO {

	return dataPickup
}
