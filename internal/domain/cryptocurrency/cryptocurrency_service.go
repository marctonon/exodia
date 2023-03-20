package cryptocurrency

type CryptocurrencyService interface {
	Sell(crypto *Cryptocurrency, amount float64) float64
	GetBalance(crypto *Cryptocurrency) float64
}

type cryptocurrencyService struct{}

func NewCryptocurrencyService() CryptocurrencyService {
	return &cryptocurrencyService{}
}

func (cs *cryptocurrencyService) Sell(crypto *Cryptocurrency, amount float64) float64 {
	return crypto.Sell(amount)
}

func (cs *cryptocurrencyService) GetBalance(crypto *Cryptocurrency) float64 {
	return crypto.GetBalance()
}
