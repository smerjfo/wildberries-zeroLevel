package nats_streaming

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	if len(os.Getenv("STANCLUSTERID")) == 0 {
		if err := os.Setenv("STANCLUSTERID", "wildberries"); err != nil {
			return err
		}
	}
	if len(os.Getenv("CLIENTID")) == 0 {
		if err := os.Setenv("CLIENTID", "user"); err != nil {
			return err
		}
	}
	if len(os.Getenv("NATSHOST")) == 0 {
		if err := os.Setenv("NATSHOST", "localhost"); err != nil {
			return err
		}
	}
	if len(os.Getenv("NATSPORT")) == 0 {
		if err := os.Setenv("NATSPORT", "4223"); err != nil {
			return err
		}
	}
	return nil
}

type Conn struct {
	Connection stan.Conn
}

func New() (*Conn, error) {
	logrus.Infof("Trying to create connection with nats-streaming on port: %s...\n", os.Getenv("NATSPORT"))
	conn, err := stan.Connect(os.Getenv("STANCLUSTERID"), os.Getenv("CLIENTID"),
		stan.NatsURL(fmt.Sprintf("%s%s:%s", "nats://", os.Getenv("NATSHOST"), os.Getenv("NATSPORT"))),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(
			func(_ stan.Conn, err error) {
				log.Fatalf("Connection lost: %s", err)
			}))
	if err != nil {
		logrus.Errorf("NATS connection error: %s ", err)
		return nil, err
	}
	logrus.Infoln("Successfully connected.")
	return &Conn{
		Connection: conn,
	}, nil
}
