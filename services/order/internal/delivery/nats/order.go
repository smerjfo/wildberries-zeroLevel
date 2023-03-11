package nats

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	nats "l0/services/order/internal/delivery/nats/order"
)

func (d *Delivery) handelFunc(msg *stan.Msg) {
	var order nats.Order
	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		logrus.Errorln("Error parsing message data:", err)
		return
	}
	fmt.Println("Received person: ", order)
	domain, err := d.toDomainOrder(&order)
	if err != nil {
		logrus.Errorln("Cannot parse received object to domain entity: ", err)
		return
	}
	response, err := d.uc.Create(domain)
	if err != nil {
		logrus.Errorln("Error while creating object: ", err)
		return
	}
	logrus.Infoln("Successfully created object: ")
	response.Print()
}
