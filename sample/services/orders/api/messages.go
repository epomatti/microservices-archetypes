package api

type OrderRequestBody struct {
	Quantity uint
	Delivery Delivery
}

type Delivery struct {
	Address string
}
