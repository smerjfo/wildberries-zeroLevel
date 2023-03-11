package http

import (
	"l0/services/order/internal/delivery/http/order"
	"l0/services/order/internal/domain"
)

func (d *Delivery) toOrderDAO(orderDom *domain.Order) *order.Order {
	return &order.Order{
		OrderUID:          orderDom.OrderUID.String(),
		TrackNumber:       orderDom.TrackNumber.String(),
		Entry:             orderDom.Entry.String(),
		Delivery:          *d.toDeliveryDAO(orderDom),
		Payment:           *d.toPaymentDAO(orderDom),
		Items:             d.toItemsResp(orderDom),
		Locale:            orderDom.Locale.String(),
		InternalSignature: orderDom.InternalSignature.String(),
		CustomerID:        orderDom.CustomerID.String(),
		DeliveryService:   orderDom.DeliveryService.String(),
		ShardKey:          orderDom.ShardKey.String(),
		SmID:              orderDom.SmID.Int(),
		DateCreated:       orderDom.DateCreated.Time(),
		OffShard:          orderDom.OffShard.String(),
	}
}

func (d *Delivery) toDeliveryDAO(delivery *domain.Order) *order.Delivery {
	return &order.Delivery{
		Name:    delivery.Delivery.Name.String(),
		Phone:   delivery.Delivery.Phone.String(),
		Zip:     delivery.Delivery.Zip.String(),
		City:    delivery.Delivery.City.String(),
		Address: delivery.Delivery.Address.String(),
		Region:  delivery.Delivery.Region.String(),
		Email:   delivery.Delivery.Email.String(),
	}
}

func (d *Delivery) toPaymentDAO(payment *domain.Order) *order.Payment {
	return &order.Payment{
		Transaction:  payment.Payment.Transaction.String(),
		RequestID:    payment.Payment.RequestID.String(),
		Currency:     payment.Payment.Currency.String(),
		Provider:     payment.Payment.Provider.String(),
		Amount:       payment.Payment.Amount.Int(),
		PaymentDT:    payment.Payment.PaymentDT.String(),
		Bank:         payment.Payment.Bank.String(),
		DeliveryCost: payment.Payment.DeliveryCost.Int(),
		GoodsTotal:   payment.Payment.GoodsTotal.Int(),
		CustomFee:    payment.Payment.CustomFee.Int(),
	}
}
func (d *Delivery) toItemsResp(items *domain.Order) []order.Item {
	daoItems := make([]order.Item, 0)
	for i, _ := range items.Items {
		result := d.toItemResp(&items.Items[i])
		daoItems = append(daoItems, result)
	}
	return daoItems
}

func (d *Delivery) toItemResp(item *domain.Item) order.Item {
	return order.Item{
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
