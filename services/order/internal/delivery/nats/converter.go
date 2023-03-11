package nats

import (
	"fmt"
	orderType "l0/services/order/internal/delivery/nats/order"
	"l0/services/order/internal/domain"
	"strconv"
)

func (d *Delivery) toDomainOrder(order *orderType.Order) (*domain.Order, error) {
	delivery := &order.Delivery
	payment := &order.Payment
	items := order.Items

	newDelivery, err := domain.NewDelivery(
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create delivery object %w", err)
	}

	newPayment, err := domain.NewPayment(
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		strconv.Itoa(payment.Amount),
		payment.PaymentDT,
		payment.Bank,
		strconv.Itoa(payment.DeliveryCost),
		strconv.Itoa(payment.GoodsTotal),
		strconv.Itoa(payment.CustomFee),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create payment object %w", err)
	}
	var domainItems []domain.Item
	for i, item := range items {
		domainItem, err := domain.NewItem(
			strconv.Itoa(item.ChrtID),
			item.TrackNumber,
			strconv.Itoa(item.Price),
			item.RID,
			item.Name,
			strconv.Itoa(item.Sale),
			item.Size,
			strconv.Itoa(item.TotalPrice),
			strconv.Itoa(item.NmID),
			item.Brand,
			strconv.Itoa(item.Status))
		if err != nil {
			return nil, fmt.Errorf("cannot create item object #%d  %w", i, err)
		}
		domainItems = append(domainItems, domainItem)
	}

	newOrder, err := domain.New(
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		newDelivery,
		newPayment,
		domainItems,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		strconv.Itoa(order.SmID),
		order.DateCreated.String(),
		order.OffShard)
	if err != nil {
		return nil, fmt.Errorf("cannot create order object: %w", err)
	}
	return &newOrder, nil
}
