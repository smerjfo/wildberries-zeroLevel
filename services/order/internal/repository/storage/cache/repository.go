package cache

import (
	"github.com/sirupsen/logrus"
	"l0/services/order/internal/domain"
)

var repo *repository

type repository struct {
	cache map[string]domain.Order
}

func GetInstance() *repository {
	if repo == nil {
		logrus.Infoln("Initialization of the cache...")
		repo = &repository{
			cache: make(map[string]domain.Order),
		}
	}
	return repo
}
