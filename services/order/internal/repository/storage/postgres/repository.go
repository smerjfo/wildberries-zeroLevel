package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MIGRATIONS_DIR", "./services/order/internal/migrations")
	viper.SetDefault("ORDERLIMIT", "1000")
}

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (*Repository, error) {
	logrus.Infoln("Initialization of the repository...")
	if err := migrations(db); err != nil {
		return nil, err
	}
	var r = &Repository{
		db: db,
	}
	logrus.Infoln("Successfully")
	return r, nil
}

func migrations(pool *pgxpool.Pool) error {
	db, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		logrus.Errorf("Error occurs while initialization of repository")
		return err
	}
	defer func() {
		if errClose := db.Close(); errClose != nil {
			err = errClose
			return
		}
	}()
	dir := viper.GetString("MIGRATIONS_DIR")
	goose.SetTableName("order_v1")
	logrus.Infoln("Starting migrations")
	if err = goose.Run("up", db, dir); err != nil {
		logrus.Errorf("Error of migrations: %s\n", err)
		return fmt.Errorf("goose up error: %w", err)
	}
	logrus.Infoln("Migrations done")
	return nil
}
