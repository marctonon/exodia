package card

type CardService interface {
	LoadBalance(card *Card, amount float64)
	AuthorizePurchase(card *Card, amount float64) bool
}

type cardService struct{}

func NewCardService() CardService {
	return &cardService{}
}

func (cs *cardService) LoadBalance(card *Card, amount float64) {
	card.LoadBalance(amount)
}

func (cs *cardService) AuthorizePurchase(card *Card, amount float64) bool {
	return card.AuthorizePurchase(amount)
}
