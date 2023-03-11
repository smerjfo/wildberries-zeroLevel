package domain

import (
	"fmt"
	"l0/pkg/email"
	"l0/pkg/phone"
	"l0/services/order/internal/domain/anyString"
)

type Delivery struct {
	Name    anyString.AnyString
	Phone   phone.Phone
	Zip     anyString.AnyString
	City    anyString.AnyString
	Address anyString.AnyString
	Region  anyString.AnyString
	Email   email.Email
}

func (d Delivery) print() {
	fmt.Println("	Name: " + d.Name)
	fmt.Println("	Phone: " + d.Phone.String())
	fmt.Println("	Zip: " + d.Zip)
	fmt.Println("	City: " + d.City)
	fmt.Println("	Address: " + d.Address)
	fmt.Println("	Region: " + d.Region)
	fmt.Println("	Email: " + d.Email.String())
}

func NewDelivery(
	name string,
	phoneNum string,
	zip string,
	city string,
	address string,
	region string,
	mail string) (Delivery, error) {
	n, err := anyString.New(name)
	if err != nil {
		return Delivery{}, err
	}
	p, err := phone.New(phoneNum)
	if err != nil {
		return Delivery{}, err
	}
	z, err := anyString.New(zip)
	if err != nil {
		return Delivery{}, err
	}
	c, err := anyString.New(city)
	if err != nil {
		return Delivery{}, err
	}
	a, err := anyString.New(address)
	if err != nil {
		return Delivery{}, err
	}
	r, err := anyString.New(region)
	if err != nil {
		return Delivery{}, err
	}
	m, err := email.New(mail)
	if err != nil {
		return Delivery{}, err
	}
	return Delivery{
		Name:    n,
		Phone:   p,
		Zip:     z,
		City:    c,
		Address: a,
		Region:  r,
		Email:   m,
	}, nil
}
