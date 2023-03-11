package order

import (
	"l0/services/order/internal/useCase/adapters/storage"
)

type UseCase struct {
	adapterStorage storage.Order
	adapterCache   storage.Order
}

func New(storage storage.Order, cache storage.Order) *UseCase {
	var uc = &UseCase{adapterStorage: storage, adapterCache: cache}
	return uc
}
