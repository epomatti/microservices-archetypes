package database

// Create will insert a new delivery
func Create(address string, orderID uint) (Delivery, error) {
	delivery := Delivery{
		Address: address,
		OrderID: orderID,
	}
	return delivery, db.Create(&delivery).Error
}

// FindbyId will select a delivery by it's primary key
func FindById(id uint) (Delivery, error) {
	var delivery Delivery
	err := db.First(&delivery, id).Error
	return delivery, err
}
