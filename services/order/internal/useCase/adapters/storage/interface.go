package storage

import "l0/services/order/internal/domain"

type Order interface {
	Create(orders *domain.Order) (*domain.Order, error)
	OrderReader
}
type OrderReader interface {
	ReadByID(ID string) (*domain.Order, error)
}
