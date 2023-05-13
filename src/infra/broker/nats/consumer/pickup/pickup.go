package mail

import (
	"encoding/json"
	"log"

	dto "books/src/app/dto/pickup"
	natsBroker "books/src/infra/broker/nats"
	booksConst "books/src/infra/constants"

	useCase "books/src/app/usecase/pickup"

	"github.com/nats-io/nats.go"
)

const (
	FcmMaxRecipient    = 1000
	FcmMaxRecipientSDK = 500
	PriorityHigh       = "high"
)

type NotifPickupInterface interface {
	InitNats()
}

type PickupWorkerImpl struct {
	nats     *natsBroker.Nats
	subjects string
	queues   string
	UseCase  useCase.PickUpUCInterface
}

func NewPickUpWorker(
	Nats *natsBroker.Nats, useCase useCase.PickUpUCInterface,
) NotifPickupInterface {

	subjects := booksConst.BOOKS

	queues := booksConst.BOOKS_QUEUE

	pickUpWorkerImpl := &PickupWorkerImpl{
		nats:     Nats,
		subjects: subjects,
		queues:   queues,
		UseCase:  useCase,
	}

	if Nats.Status {
		pickUpWorkerImpl.InitNats()
	}

	return pickUpWorkerImpl
}

func (p *PickupWorkerImpl) InitNats() {

	go eventNotificationWorker(p)

}

func eventNotificationWorker(t *PickupWorkerImpl) {

	_, err := t.nats.Conn.QueueSubscribe(t.subjects, t.queues, func(msg *nats.Msg) {

		dataConsume := dto.ReqPickupDTO{}
		err := json.Unmarshal(msg.Data, &dataConsume)
		if err != nil {
			log.Printf("%+v", err)

			return
		}

		t.UseCase.AddDataPickUp(&dataConsume)
		if err != nil {
			log.Printf("%+v", err)

			return
		}
	})

	if err != nil {
		log.Fatal(err)
	}
	t.nats.Conn.Flush()
	if err := t.nats.Conn.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", t.subjects)

}
