package postgres

import (
	"fmt"
	"l0/services/order/internal/domain"
	"l0/services/order/internal/repository/storage/postgres/dao"
	"strconv"
)

func (r *Repository) toDomainOrder(order *dao.Order, delivery *dao.Delivery, payment *dao.Payment, items []dao.Item) (*domain.Order, error) {

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
		payment.PaymentDT.String(),
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

func (r *Repository) toOrderDAO(order *domain.Order) *dao.Order {
	return &dao.Order{
		OrderUID:          order.OrderUID.String(),
		TrackNumber:       order.TrackNumber.String(),
		Entry:             order.Entry.String(),
		Locale:            order.Locale.String(),
		InternalSignature: order.InternalSignature.String(),
		CustomerID:        order.CustomerID.String(),
		DeliveryService:   order.DeliveryService.String(),
		ShardKey:          order.ShardKey.String(),
		SmID:              order.SmID.Int(),
		DateCreated:       order.DateCreated.Time(),
		OffShard:          order.OffShard.String(),
	}
}

func (r *Repository) toDeliveryDAO(delivery *domain.Order) *dao.Delivery {
	return &dao.Delivery{
		Name:    delivery.Delivery.Name.String(),
		Phone:   delivery.Delivery.Phone.String(),
		Zip:     delivery.Delivery.Zip.String(),
		City:    delivery.Delivery.City.String(),
		Address: delivery.Delivery.Address.String(),
		Region:  delivery.Delivery.Region.String(),
		Email:   delivery.Delivery.Email.String(),
	}
}

func (r *Repository) toPaymentDAO(payment *domain.Order) *dao.Payment {
	return &dao.Payment{
		Transaction:  payment.Payment.Transaction.String(),
		RequestID:    payment.Payment.RequestID.String(),
		Currency:     payment.Payment.Currency.String(),
		Provider:     payment.Payment.Provider.String(),
		Amount:       payment.Payment.Amount.Int(),
		PaymentDT:    payment.Payment.PaymentDT.Time(),
		Bank:         payment.Payment.Bank.String(),
		DeliveryCost: payment.Payment.DeliveryCost.Int(),
		GoodsTotal:   payment.Payment.GoodsTotal.Int(),
		CustomFee:    payment.Payment.CustomFee.Int(),
	}
}

func (r *Repository) toItemDAO(item *domain.Item) *dao.Item {
	return &dao.Item{
		ChrtID:      item.ChrtID.Int(),
		TrackNumber: item.TrackNumber.String(),
		Price:       item.Price.Int(),
		RID:         item.RID.String(),
		Name:        item.Name.String(),
		Sale:        item.Sale.Int(),
		Size:        item.Size.String(),
		TotalPrice:  item.TotalPrice.Int(),
		NmID:        item.NmID.Int(),
		Brand:       item.Brand.String(),
		Status:      item.Status.Int(),
	}
}

func (r *Repository) toItemsDAO(items *domain.Order) []*dao.Item {
	daoItems := make([]*dao.Item, 0)
	for i, _ := range items.Items {
		result := r.toItemDAO(&items.Items[i])
		daoItems = append(daoItems, result)
	}
	return daoItems
}
