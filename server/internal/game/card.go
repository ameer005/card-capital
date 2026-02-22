package game

type Card struct {
	Id                string
	Name              string
	price             float64
	quantity          int
	availableQuantity int
}

func NewCard(name string, price float64, quantity int, id string) *Card {
	return &Card{
		Id:                id,
		Name:              name,
		price:             price,
		quantity:          quantity,
		availableQuantity: quantity,
	}
}

func (c *Card) UpdatePrice(price float64) bool {
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
func (c *Card) Price() float64 {
	return c.price
}

func (c *Card) Quantity() int {
	return c.quantity
}

func (c *Card) AvailaleQuantity() int {
	return c.availableQuantity
}
