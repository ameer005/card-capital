package game

type Card struct {
	Name              string
	price             int
	quantity          int
	availableQuantity int
}

func NewCard(name string, price, quantity, availabelQuantity int) *Card {
	return &Card{
		Name:              name,
		price:             price,
		quantity:          quantity,
		availableQuantity: availabelQuantity,
	}
}

// updaters
func (c *Card) UpdatePrice(price int) bool {
	c.price = price
	return true
}

func (c *Card) DecreaseAvailableQuantity(amount int) bool {
	newQty := c.availableQuantity - amount

	if newQty < 0 {
		return false
	}

	c.availableQuantity = newQty

	return true
}

func (c *Card) IncreaseAvailableQuantity(amount int) bool {
	newQty := c.availableQuantity + amount

	if newQty > c.quantity {
		return false
	}

	c.availableQuantity = newQty

	return true
}

// getters
func (c *Card) Price() int {
	return c.price
}

func (c *Card) Quantity() int {
	return c.quantity
}

func (c *Card) AvailaleQuantity() int {
	return c.availableQuantity
}
