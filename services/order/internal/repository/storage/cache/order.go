package cache

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"l0/services/order/internal/domain"
)

func (r *repository) ReadByID(ID string) (*domain.Order, error) {
	order, exists := r.cache[ID]
	if !exists {
		return nil, fmt.Errorf("order with this ID doesnt exists in cache")
	}
	return &order, nil
}
func (r *repository) Create(order *domain.Order) (*domain.Order, error) {
	_, exists := r.cache[order.OrderUID.String()]
	if exists {
		logrus.Errorln("order with current ID already exists in cache")
		return nil, fmt.Errorf("order with current ID already exists in cache")
	}
	r.cache[order.OrderUID.String()] = *order
	return order, nil
}
func (r *repository) Initialize(orders []*domain.Order) error {
	logrus.Infoln("Filling the cache...")
	for _, order := range orders {
		_, err := r.Create(order)
		if err != nil {
			return err
		}
	}
	return nil
}
