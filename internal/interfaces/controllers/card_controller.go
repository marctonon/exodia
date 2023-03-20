package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/marctonon/exodia/internal/domain/card"
	"github.com/marctonon/exodia/internal/domain/cryptocurrency"
	"github.com/marctonon/exodia/internal/usecases"
)

type CardController interface {
	AuthorizeCreditCardPurchase(w http.ResponseWriter, r *http.Request)
}

type cardController struct {
	cardService         card.CardService
	cryptoService       cryptocurrency.CryptocurrencyService
	cardUseCase         usecases.CardUseCase
}

func NewCardController() CardController {
	cardService := card.NewCardService()
	cryptoService := cryptocurrency.NewCryptocurrencyService()
	cardUseCase := usecases.NewCardUseCase(cardService, cryptoService)

	return &cardController{
		cardService:         cardService,
		cryptoService:       cryptoService,
		cardUseCase:         cardUse
	}
}

type authorizeCreditCardPurchaseRequest struct {
	CardNumber string  `json:"card_number"`
	Amount     float64 `json:"amount"`
}

type authorizeCreditCardPurchaseResponse struct {
	Status string `json:"status"`
}

func (cc *cardController) AuthorizeCreditCardPurchase(w http.ResponseWriter, r *http.Request) {
	var request authorizeCreditCardPurchaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := cc.cardUseCase.AuthorizeCreditCardPurchase(request.CardNumber, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := authorizeCreditCardPurchaseResponse{
		Status: "approved",
	}

	json.NewEncoder(w).Encode(response)
}
