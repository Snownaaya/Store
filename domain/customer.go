package domain

type Customer struct {
	inventory
}

func newCustomer() *Customer {
	return &Customer{
		inventory: inventory{
			Money:    1000000000,
			Products: map[string](*[]Product){"Electronics": &[]Product{}},
		},
	}
}

func (c *Customer) HasEnoughMoney(price int) bool {
	return c.Money >= price
}

func (c *Customer) TakeProduct(product Product) {
	products := c.Products["Electronics"]

	*products = append((*products), product)
}

func (c *Customer) Pay(price int) {
	c.Money -= price
}
