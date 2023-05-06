package api

type DeliveryRequestBody struct {
	Address string
	OrderID uint `json:"orderId"`
}
