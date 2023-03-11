package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}
func initDefaultEnv() error {
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "localhost"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "disable"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "5431"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "wildberries"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGUSER", "user"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "password"); err != nil {
			return err
		}
	}
	return nil
}

type Store struct {
	Pool *pgxpool.Pool
}

func New() (*Store, error) {
	logrus.Infoln("Trying to create DB connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		logrus.Errorln("Error occurs: ", err)
		return nil, err
	}
	conn, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logrus.Errorln("Error occurs: ", err)
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		logrus.Errorln("Error occurs: ", err)
		return nil, err
	}
	logrus.Infoln("Successfully connected.")
	return &Store{Pool: conn}, nil

}
