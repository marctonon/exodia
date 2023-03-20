package card

type Card struct {
	ID             int
	Number         string
	Balance        float64
	Holder         string
	ExpirationDate string
	SecurityCode   string
}

func (c *Card) LoadBalance(amount float64) {
	c.Balance += amount
}

func (c *Card) AuthorizePurchase(amount float64) bool {
	if c.Balance >= amount {
		c.Balance -= amount
		return true
	}

	return false
}
