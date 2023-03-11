package domain

import (
	"fmt"
	"l0/services/order/internal/domain/anyInt"
	"l0/services/order/internal/domain/anyString"
)

type Item struct {
	ChrtID      anyInt.Int
	TrackNumber anyString.AnyString
	Price       anyInt.Int
	RID         anyString.AnyString
	Name        anyString.AnyString
	Sale        anyInt.Int
	Size        anyString.AnyString
	TotalPrice  anyInt.Int
	NmID        anyInt.Int
	Brand       anyString.AnyString
	Status      anyInt.Int
}

func (i Item) print() {
	fmt.Println("	ChrtID: " + i.ChrtID.String())
	fmt.Println("	TrackNumber: " + i.TrackNumber)
	fmt.Println("	Price: " + i.Price.String())
	fmt.Println("	RID: " + i.RID)
	fmt.Println("	Name: " + i.Name)
	fmt.Println("	Sale: " + i.Sale.String())
	fmt.Println("	Size: " + i.Size)
	fmt.Println("	TotalPrice: " + i.TotalPrice.String())
	fmt.Println("	NmID: " + i.NmID.String())
	fmt.Println("	Brand: " + i.Brand)
	fmt.Println("	Status: " + i.Status.String())
}

func NewItem(
	ChrtID string,
	TrackNumber string,
	Price string,
	RID string,
	Name string,
	Sale string,
	Size string,
	TotalPrice string,
	NmID string,
	Brand string,
	Status string) (Item, error) {
	chrt, err := anyInt.NewStr(ChrtID)
	if err != nil {
		return Item{}, err
	}
	tn, err := anyString.New(TrackNumber)
	if err != nil {
		return Item{}, err
	}
	price, err := anyInt.NewStr(Price)
	if err != nil {
		return Item{}, err
	}
	rid, err := anyString.New(RID)
	if err != nil {
		return Item{}, err
	}
	name, err := anyString.New(Name)
	if err != nil {
		return Item{}, err
	}
	size, err := anyString.New(Size)
	if err != nil {
		return Item{}, err
	}
	br, err := anyString.New(Brand)
	if err != nil {
		return Item{}, err
	}
	sale, err := anyInt.NewStr(Sale)
	if err != nil {
		return Item{}, err
	}
	nmid, err := anyInt.NewStr(NmID)
	if err != nil {
		return Item{}, err
	}
	tp, err := anyInt.NewStr(TotalPrice)
	if err != nil {
		return Item{}, err
	}
	status, err := anyInt.NewStr(Status)
	if err != nil {
		return Item{}, err
	}
	return Item{
		ChrtID:      chrt,
		TrackNumber: tn,
		Sale:        sale,
		Size:        size,
		TotalPrice:  tp,
		Status:      status,
		NmID:        nmid,
		Name:        name,
		Brand:       br,
		Price:       price,
		RID:         rid,
	}, nil

}
