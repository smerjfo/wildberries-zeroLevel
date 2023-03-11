package domain

import (
	"fmt"
	"l0/services/order/internal/domain/anyInt"
	"l0/services/order/internal/domain/anyString"
	"l0/services/order/internal/domain/timestamp"
)

//TODO create DDD

type Order struct {
	OrderUID          anyString.AnyString
	TrackNumber       anyString.AnyString
	Entry             anyString.AnyString
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            anyString.AnyString
	InternalSignature anyString.AnyString
	CustomerID        anyString.AnyString
	DeliveryService   anyString.AnyString
	ShardKey          anyString.AnyString
	SmID              anyInt.Int
	DateCreated       timestamp.Timestamp
	OffShard          anyString.AnyString
}

func (o *Order) Print() {
	fmt.Println("Order_UID: " + o.OrderUID)
	fmt.Println("TrackNumber: " + o.TrackNumber)
	fmt.Println("Entry: " + o.Entry)
	fmt.Println("DELIVERY: ")
	o.Delivery.print()
	fmt.Println("PAYMENT: ")
	o.Payment.print()
	for i, arr := range o.Items {
		fmt.Println("ITEM №" + string(i))
		arr.print()
	}
	fmt.Println("Locale: " + o.Locale)
	fmt.Println("InternalSignature: " + o.InternalSignature)
	fmt.Println("CustomerID: " + o.CustomerID)
	fmt.Println("DeliveryService: " + o.DeliveryService)
	fmt.Println("ShardKey: " + o.ShardKey)
	fmt.Println("SmID: " + o.SmID.String())
	fmt.Println("DateCreated: " + o.DateCreated.String())
	fmt.Println("OffShard: " + o.OffShard)
}

func New(
	orderId string,
	trackNum string,
	entry string,
	delivery Delivery,
	payment Payment,
	items []Item,
	locale string,
	internalSignature string,
	customerID string,
	deliveryService string,
	shardKey string,
	smID string,
	dateCreated string,
	offShard string,
) (Order, error) {
	oID, err := anyString.New(orderId)
	if err != nil {
		return Order{}, err
	}
	tN, err := anyString.New(trackNum)
	if err != nil {
		return Order{}, err
	}
	e, err := anyString.New(entry)
	if err != nil {
		return Order{}, err
	}
	l, err := anyString.New(locale)
	if err != nil {
		return Order{}, err
	}
	iS, err := anyString.New(internalSignature)
	if err != nil {
		return Order{}, err
	}
	cID, err := anyString.New(customerID)
	if err != nil {
		return Order{}, err
	}
	dS, err := anyString.New(deliveryService)
	if err != nil {
		return Order{}, err
	}
	sK, err := anyString.New(shardKey)
	if err != nil {
		return Order{}, err
	}
	oS, err := anyString.New(offShard)
	if err != nil {
		return Order{}, err
	}
	sID, err := anyInt.NewStr(smID)
	if err != nil {
		return Order{}, err
	}
	dC, err := timestamp.New(dateCreated)
	if err != nil {
		return Order{}, err
	}
	return Order{
		OrderUID:          oID,
		TrackNumber:       tN,
		Entry:             e,
		Locale:            l,
		InternalSignature: iS,
		CustomerID:        cID,
		DeliveryService:   dS,
		ShardKey:          sK,
		OffShard:          oS,
		SmID:              sID,
		DateCreated:       dC,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
	}, nil

}
