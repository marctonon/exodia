package usecases

import (
	"errors"

	"github.com/marctonon/exodia/internal/domain/card"
	"github.com/marctonon/exodia/internal/domain/cryptocurrency"
	"github.com/marctonon/exodia/internal/interfaces/repositories"
)

type CardUseCase interface {
	AuthorizeCreditCardPurchase(cardNumber string, amount float64) (*card.Card, error)
}

type cardUseCase struct {
	cardService         card.CardService
	cryptoRepository     repositories.CryptocurrencyRepository
}

func NewCardUseCase(cardService card.CardService, cryptoService cryptocurrency.CryptocurrencyService) CardUseCase {
	return &cardUseCase{
		cardService:         cardService,
		cryptoRepository:     repositories.NewCryptocurrencyRepository(db),
	}
}

func (uc *cardUseCase) AuthorizeCreditCardPurchase(cardNumber string, amount float64) (*card.Card, error) {
	c, err := uc.getCardByNumber(cardNumber)
	if err != nil {
		return nil, err
	}

	if !uc.cardService.AuthorizePurchase(c, amount) {
		return nil, errors.New("insufficient balance")
	}

	symbol := "BTC" // Supondo que o cartão só possa ser utilizado para compra de Bitcoin

	balance, err := uc.cryptoRepository.GetBalance(symbol)
	if err != nil {
		return nil, err
	}

	if balance < amount {
		return nil, errors.New("insufficient cryptocurrency balance")
	}

	uc.cryptoRepository.UpdateBalance(symbol, balance - amount)
	uc.cardService.LoadBalance(c, amount)

	return c,
}

func (uc *cardUseCase) getCardByNumber(cardNumber string) (*card.Card, error) {
	// Consulta o cartão no banco de dados ou em outra fonte de dados
	// Suponha que o cartão seja encontrado com ID 1
	return &card.Card{
		ID:      1,
		Number:  cardNumber,
		Balance: 1000,
		Holder:  "John Doe",
		ExpirationDate: "12/25",
		SecurityCode: "123",
	}, nil
}
