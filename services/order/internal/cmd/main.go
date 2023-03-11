package main

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"l0/pkg/messageBroker/nats-streaming"
	"l0/pkg/store/postgres"
	"l0/services/order/internal/delivery/http"
	"l0/services/order/internal/delivery/nats"
	"l0/services/order/internal/repository/storage/cache"
	postgres2 "l0/services/order/internal/repository/storage/postgres"
	"l0/services/order/internal/useCase/order"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	cn, err := nats_streaming.New()
	if err != nil {

		panic(err)
	}
	defer func(cn stan.Conn) {
		logrus.Infoln("Closing nats connection ...")
		err := cn.Close()
		if err != nil {
			panic(err)
		}
	}(cn.Connection)

	err = publishMessages(cn)
	if err != nil {
		panic(err)
	}

	conn, err := postgres.New()
	if err != nil {
		logrus.Errorf("Error occurs while creaing DB connection: %s", err)
		panic(err)
	}

	defer conn.Pool.Close()

	repoStorage, err := postgres2.New(conn.Pool)
	if err != nil {
		panic(err)
	}

	cacheInst := cache.GetInstance()
	orders, err := repoStorage.ReadRowsByLimit()
	if err != nil {
		panic(err)
	}
	err = cacheInst.Initialize(orders)
	if err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	var (
		ucOrder      = order.New(repoStorage, cacheInst)
		deliveryHTTP = http.New(ucOrder)
		deliveryNATS = nats.New(ucOrder, cn)
	)
	go func() {

		err := deliveryNATS.Run(stopChan)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		err := deliveryHTTP.Run()
		if err != nil {
			panic(err)
		}
	}()

	<-stopChan
}
