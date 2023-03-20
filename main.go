package main

import (
	"log"
	"net/http"

	"github.com/marctonon/exodia/internal/interfaces/controllers"
)

func main() {
	cardController := controllers.NewCardController()

	http.HandleFunc("/authorize-credit-card-purchase", cardController.AuthorizeCreditCardPurchase)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
