package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"l0/services/order/internal/domain"
	"l0/services/order/internal/repository/storage/postgres/dao"
)

func (r *Repository) ReadByID(ID string) (*domain.Order, error) {
	if len(ID) == 0 {
		return nil, nil
	}
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		if err != nil {
			if rollback := tx.Rollback(ctx); rollback != nil {
				logrus.Infoln("Error in rollback")
				err = rollback
			}
		} else if commitErr := tx.Commit(ctx); commitErr != nil {
			logrus.Infoln("Error in commit")
			err = commitErr
		}
	}(ctx, tx)
	response, err := r.getOrderByIDTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Repository) getOrderByIDTx(ctx context.Context, tx pgx.Tx, ID string) (*domain.Order, error) {

	order := dao.Order{}
	orderQuery := "SELECT order_uid, track_number, entry, locale, internal_signature, customer_id,delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders WHERE order_uid = $1"
	err := tx.QueryRow(ctx, orderQuery, ID).Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature,
		&order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SmID, &order.DateCreated, &order.OffShard)
	if err != nil {
		return nil, fmt.Errorf("failed to get order details: %w", err)
	}
	delivery, err := r.readDeliveryByID(ctx, tx, ID)
	if err != nil {
		return nil, err
	}
	payment, err := r.readPaymentByID(ctx, tx, ID)
	if err != nil {
		return nil, err
	}
	items, err := r.readItemsByID(ctx, tx, ID)
	if err != nil {
		return nil, err
	}

	domainOrder, err := r.toDomainOrder(&order, delivery, payment, items)
	if err != nil {
		return nil, fmt.Errorf("failed to convert order from dao to domain details: %w", err)
	}
	return domainOrder, err
}

func (r *Repository) readDeliveryByID(ctx context.Context, tx pgx.Tx, ID string) (*dao.Delivery, error) {
	delivery := dao.Delivery{}
	deliveryQuery := "SELECT name, phone, zip, city, address, region, email FROM order_deliveries WHERE order_uid = $1"
	err := tx.QueryRow(ctx, deliveryQuery, ID).Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery details: %w", err)
	}
	return &delivery, nil
}

func (r *Repository) readPaymentByID(ctx context.Context, tx pgx.Tx, ID string) (*dao.Payment, error) {
	payment := dao.Payment{}
	paymentQuery := "SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM order_payments WHERE order_uid = $1"
	err := tx.QueryRow(ctx, paymentQuery, ID).Scan(&payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment details: %w", err)
	}
	return &payment, nil
}

func (r *Repository) readItemsByID(ctx context.Context, tx pgx.Tx, ID string) ([]dao.Item, error) {
	items := make([]dao.Item, 0)
	itemQuery := "SELECT i.chrt_id, i.track_number, i.price, i.rid, i.name, i.sale, i.size, i.total_price, i.nm_id, i.brand, i.status FROM items i JOIN order_items oi ON i.chrt_id = oi.chrt_id WHERE oi.order_uid=$1"
	rows, err := tx.Query(ctx, itemQuery, ID)
	for rows.Next() {
		item := dao.Item{}
		err = rows.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to get item details: %w", err)
		}
		items = append(items, item)
	}
	return items, err
}

func (r *Repository) ReadRowsByLimit() ([]*domain.Order, error) {
	logrus.Infoln("Selecting last orders...")
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		if err != nil {
			if rollback := tx.Rollback(ctx); rollback != nil {
				logrus.Infoln("Error in rollback")
				err = rollback
			}
		} else if commitErr := tx.Commit(ctx); commitErr != nil {
			logrus.Infoln("Error in commit")
			err = commitErr
		}
	}(ctx, tx)
	response, err := r.getOrdersByLimitTx(ctx, tx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Repository) getOrdersByLimitTx(ctx context.Context, tx pgx.Tx) ([]*domain.Order, error) {
	orders := make([]*domain.Order, 0)
	orderQuery := "SELECT order_uid, track_number, entry, locale, internal_signature, customer_id,delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders LIMIT $1"
	rows, err := tx.Query(ctx, orderQuery, viper.GetInt("ORDERSLIMIT"))
	for rows.Next() {
		order := dao.Order{}
		err = rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature,
			&order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SmID, &order.DateCreated, &order.OffShard)
		if err != nil {
			return nil, err
		}
		delivery, err := r.readDeliveryByID(ctx, tx, order.OrderUID)
		if err != nil {
			return nil, err
		}
		payment, err := r.readPaymentByID(ctx, tx, order.OrderUID)
		if err != nil {
			return nil, err
		}
		items, err := r.readItemsByID(ctx, tx, order.OrderUID)
		if err != nil {
			return nil, err
		}

		domainOrder, err := r.toDomainOrder(&order, delivery, payment, items)
		orders = append(orders, domainOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to convert order from dao to domain details: %w", err)
		}

	}
	return orders, nil
}

func (r *Repository) Create(order *domain.Order) (*domain.Order, error) {
	if order == nil {
		return nil, nil
	}
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		logrus.Infoln("ERROR")
		return nil, err
	}
	defer func(ctx context.Context, tx pgx.Tx) {
		if err != nil {
			if rollback := tx.Rollback(ctx); rollback != nil {
				logrus.Infoln("Error in rollback")
				err = rollback
			}
		} else if commitErr := tx.Commit(ctx); commitErr != nil {
			logrus.Infoln("Error in commit")
			err = commitErr
		}
	}(ctx, tx)
	response, err := r.createOrderTx(ctx, tx, order)

	if err != nil {
		return nil, err
	}
	logrus.Infoln("SUCCESSFULLY CREATED")
	return response, nil
}

func (r *Repository) createOrderTx(ctx context.Context, tx pgx.Tx, order *domain.Order) (*domain.Order, error) {
	logrus.Infoln("CREATING OBJECT...")
	var exists bool
	err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT order_uid FROM orders WHERE order_uid=$1)", order.OrderUID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("order with current uid already exists")
	}

	query := "INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	daoOrder := r.toOrderDAO(order)

	_, err = tx.Exec(ctx, query, &daoOrder.OrderUID, &daoOrder.TrackNumber, &daoOrder.Entry, &daoOrder.Locale, &daoOrder.InternalSignature,
		&daoOrder.CustomerID, &daoOrder.DeliveryService, &daoOrder.ShardKey, &daoOrder.SmID, &daoOrder.DateCreated, &daoOrder.OffShard)
	if err != nil {
		return nil, err
	}

	query = "INSERT INTO order_deliveries (order_uid, name, phone, zip, city, address, region, email) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	daoDelivery := r.toDeliveryDAO(order)
	_, err = tx.Exec(ctx, query, &daoOrder.OrderUID, &daoDelivery.Name, &daoDelivery.Phone, &daoDelivery.Zip, &daoDelivery.City, &daoDelivery.Address, &daoDelivery.Region, &daoDelivery.Email)
	if err != nil {
		return nil, err
	}

	query = "INSERT INTO order_payments (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	daoPayment := r.toPaymentDAO(order)
	_, err = tx.Exec(ctx, query, &daoOrder.OrderUID, &daoPayment.Transaction, &daoPayment.RequestID, &daoPayment.Currency, &daoPayment.Provider, &daoPayment.Amount, &daoPayment.PaymentDT,
		&daoPayment.Bank, &daoPayment.DeliveryCost, &daoPayment.GoodsTotal, &daoPayment.CustomFee)
	if err != nil {
		return nil, err
	}
	_, err = r.createItems(ctx, tx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) createItems(ctx context.Context, tx pgx.Tx, order *domain.Order) (*domain.Order, error) {
	items := r.toItemsDAO(order)
	for _, item := range items {
		var exists bool
		err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT chrt_id FROM items WHERE chrt_id=$1)", item.ChrtID).Scan(&exists)
		if err != nil {
			return nil, err
		}
		query := "INSERT INTO order_items (order_uid, chrt_id) VALUES ($1,$2)"
		if exists {
			_, err = tx.Exec(ctx, query, &order.OrderUID, &item.ChrtID)
			if err != nil {
				return nil, err
			}
			continue
		}
		itemsQ := "INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) " +
			"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
		_, err = tx.Exec(ctx, itemsQ, &item.ChrtID, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return nil, err
		}
		_, err = tx.Exec(ctx, query, &order.OrderUID, &item.ChrtID)
		if err != nil {
			return nil, err
		}
	}
	return order, nil
}
