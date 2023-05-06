package database

func Create(address, orderId string) (id string, err error) {
	delivery := Deliveries{
		Address: address,
		OrderId: orderId,
	}
	return delivery.Id, db.Create(&delivery).Error
}
