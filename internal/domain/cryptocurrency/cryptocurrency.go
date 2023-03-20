package cryptocurrency

type Cryptocurrency struct {
	ID     int
	Symbol string
	Name   string
	Amount float64
}

func (c *Cryptocurrency) Sell(amount float64) float64 {
	if c.Amount >= amount {
		c.Amount -= amount
		return amount
	}

	return 0
}

func (c *Cryptocurrency) GetBalance() float64 {
	return c.Amount
}
