package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	nats_streaming "l0/pkg/messageBroker/nats-streaming"
	nats "l0/services/order/internal/delivery/http/order"
	"time"
)

func publishMessages(conn *nats_streaming.Conn) error {
	logrus.Infof("Trying to put objects in nats channel...")
	order := nats.Order{
		OrderUID:    "12345678",
		TrackNumber: "ABCD1234",
		Entry:       "Web",
		Delivery: nats.Delivery{
			Name:    "John Doe",
			Phone:   "555-1234",
			Zip:     "12345",
			City:    "New York",
			Address: "123 Main St",
			Region:  "NY",
			Email:   "johndoe@example.com",
		},
		Payment: nats.Payment{
			Transaction:  "123456",
			RequestID:    "987654",
			Currency:     "USD",
			Provider:     "PayPal",
			Amount:       1000,
			PaymentDT:    "2023-03-10 17:09:25.018944835 +0300 MSK",
			Bank:         "Bank of America",
			DeliveryCost: 100,
			GoodsTotal:   900,
			CustomFee:    0,
		},
		Items: []nats.Item{
			{
				ChrtID:      1,
				TrackNumber: "ABCD1234",
				Price:       500,
				RID:         "RID-001",
				Name:        "Product A",
				Sale:        0,
				Size:        "M",
				TotalPrice:  500,
				NmID:        1,
				Brand:       "Brand A",
				Status:      1,
			},
			{
				ChrtID:      2,
				TrackNumber: "ABCD1234",
				Price:       400,
				RID:         "RID-002",
				Name:        "Product B",
				Sale:        0,
				Size:        "L",
				TotalPrice:  400,
				NmID:        2,
				Brand:       "Brand B",
				Status:      1,
			},
		},
		Locale:            "en-US",
		InternalSignature: "SIG-123",
		CustomerID:        "CUST-001",
		DeliveryService:   "UPS",
		ShardKey:          "shard1",
		SmID:              123,
		DateCreated:       time.Now(),
		OffShard:          "123",
	}
	bytes, err := json.Marshal(order)
	if err != nil {
		logrus.Infoln("Error in marshalling object order: ", err)
		return err
	}
	err = conn.Connection.Publish("order", bytes)
	if err != nil {
		logrus.Infoln("Error while trying publish object in \"order\" channel: ", err)
		return err
	}
	logrus.Infoln("Successfully")
	return nil
}
