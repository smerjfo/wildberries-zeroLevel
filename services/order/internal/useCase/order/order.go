package order

import (
	"github.com/sirupsen/logrus"
	"l0/services/order/internal/domain"
)

func (uc *UseCase) Create(order *domain.Order) (*domain.Order, error) {
	_, err := uc.adapterCache.Create(order)
	if err != nil {
		logrus.Infoln("Error adding object into cache")
		return nil, err
	}
	return uc.adapterStorage.Create(order)
}

func (uc *UseCase) ReadByID(ID string) (*domain.Order, error) {
	order, err := uc.adapterCache.ReadByID(ID)
	if err != nil {
		orderStorage, err2 := uc.adapterStorage.ReadByID(ID)
		if err2 != nil {
			return nil, err2
		}
		return orderStorage, nil
	}
	return order, nil
}
