package nats_publisher

/*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : books
 */

import (
	"log"

	"books/src/infra/broker/nats"
)

type PublisherInterface interface {
	Nats(data []byte, subject string) error
}

type PushWorkerImpl struct {
	nats *nats.Nats
}

func NewPushWorker(Nats *nats.Nats) PublisherInterface {

	pushWorkerImpl := &PushWorkerImpl{
		nats: Nats,
	}

	return pushWorkerImpl
}

func (p *PushWorkerImpl) Nats(data []byte, subject string) error {
	err := p.nats.Conn.Publish(subject, data)
	if err != nil {
		return err
	}
	err = p.nats.Conn.Flush()
	if err != nil {
		return err
	}
	if err := p.nats.Conn.LastError(); err != nil {
		return err
	}

	log.Printf("Published to [%s]\n", subject)

	return nil
}
