package order

import "time"

type Order struct {
	OrderUID          string    `json:"order_UID"`
	TrackNumber       string    `json:"trackNumber"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internalSignature"`
	CustomerID        string    `json:"customerID"`
	DeliveryService   string    `json:"deliveryService"`
	ShardKey          string    `json:"shardKey"`
	SmID              int       `json:"smID"`
	DateCreated       time.Time `json:"dateCreated"`
	OffShard          string    `json:"offShard"`
}

type Item struct {
	ChrtID      int    `json:"chrtID"`
	TrackNumber string `json:"trackNumber"`
	Price       int    `json:"price"`
	RID         string `json:"RID"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"totalPrice"`
	NmID        int    `json:"nmID"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"requestID"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    string `json:"paymentDT"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"deliveryCost"`
	GoodsTotal   int    `json:"goodsTotal"`
	CustomFee    int    `json:"customFee"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
