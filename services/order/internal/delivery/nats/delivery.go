package nats

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	nats_streaming "l0/pkg/messageBroker/nats-streaming"
	"l0/services/order/internal/useCase"
	"os"
)

type Delivery struct {
	uc   useCase.Order
	conn *nats_streaming.Conn
}

func New(uc useCase.Order, conn *nats_streaming.Conn) *Delivery {
	return &Delivery{
		uc:   uc,
		conn: conn,
	}
}

func (d *Delivery) Run(stopChan <-chan os.Signal) error {

	sub, err := d.conn.Connection.Subscribe("order", d.handelFunc, stan.DurableName("my-durable"), stan.StartWithLastReceived())
	if err != nil {
		return err
	}
	logrus.Infoln("Subscribing started successfully.")

	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
		}
	}(sub)

	select {
	case <-stopChan:
		logrus.Infoln("Stop listening messages...")
		return nil
	}
}
