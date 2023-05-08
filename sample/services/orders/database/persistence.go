package database

// Create will insert a new order
func Create(quantity uint) (Order, error) {
	order := Order{
		Quantity: quantity,
	}
	return order, db.Create(&order).Error
}

// FindById will select a order by it's primary key
func FindById(id uint) (Order, error) {
	var order Order
	err := db.First(&order, id).Error
	return order, err
}
