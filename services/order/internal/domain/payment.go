package domain

import (
	"fmt"
	"l0/services/order/internal/domain/anyInt"
	"l0/services/order/internal/domain/anyString"
	"l0/services/order/internal/domain/timestamp"
)

type Payment struct {
	Transaction  anyString.AnyString
	RequestID    anyString.AnyString
	Currency     anyString.AnyString
	Provider     anyString.AnyString
	Amount       anyInt.Int
	PaymentDT    timestamp.Timestamp
	Bank         anyString.AnyString
	DeliveryCost anyInt.Int
	GoodsTotal   anyInt.Int
	CustomFee    anyInt.Int
}

func (receiver *Payment) print() {
	fmt.Println("	Transaction: " + receiver.Transaction)
	fmt.Println("	RequestID: " + receiver.RequestID)
	fmt.Println("	Currency: " + receiver.Currency)
	fmt.Println("	Provider: " + receiver.Provider)
	fmt.Println("	Amount: " + receiver.Amount.String())
	fmt.Println("	PaymentDT: " + receiver.PaymentDT.String())
	fmt.Println("	Bank: " + receiver.Bank)
	fmt.Println("	DeliveryCost: " + receiver.DeliveryCost.String())
	fmt.Println("	GoodsTotal: " + receiver.GoodsTotal.String())
	fmt.Println("	CustomFee: " + receiver.CustomFee.String())
}

func NewPayment(
	transaction string,
	requestID string,
	currency string,
	provider string,
	amount string,
	paymentDT string,
	bank string,
	deliveryCost string,
	goodsTotal string,
	customFee string) (Payment, error) {
	tr, err := anyString.New(transaction)
	if err != nil {
		return Payment{}, err
	}
	b, err := anyString.New(bank)
	if err != nil {
		return Payment{}, err
	}
	rID, err := anyString.New(requestID)
	if err != nil {
		return Payment{}, err
	}
	c, err := anyString.New(currency)
	if err != nil {
		return Payment{}, err
	}
	p, err := anyString.New(provider)
	if err != nil {
		return Payment{}, err
	}
	am, err := anyInt.NewStr(amount)
	if err != nil {
		return Payment{}, err
	}
	dC, err := anyInt.NewStr(deliveryCost)
	if err != nil {
		return Payment{}, err
	}
	gT, err := anyInt.NewStr(goodsTotal)
	if err != nil {
		return Payment{}, err
	}
	cF, err := anyInt.NewStr(customFee)
	if err != nil {
		return Payment{}, err
	}
	pDt, err := timestamp.New(paymentDT)
	if err != nil {
		return Payment{}, err
	}
	return Payment{
		Transaction:  tr,
		Bank:         b,
		RequestID:    rID,
		Currency:     c,
		CustomFee:    cF,
		Amount:       am,
		GoodsTotal:   gT,
		DeliveryCost: dC,
		Provider:     p,
		PaymentDT:    pDt,
	}, nil

}
