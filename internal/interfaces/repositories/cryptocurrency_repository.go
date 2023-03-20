package repositories

import (
	"database/sql"
)

type CryptocurrencyRepository interface {
	GetBalance(symbol string) (float64, error)
	UpdateBalance(symbol string, amount float64) error
}

type cryptocurrencyRepository struct {
	db *sql.DB
}

func NewCryptocurrencyRepository(db *sql.DB) CryptocurrencyRepository {
	return &cryptocurrencyRepository{
		db: db,
	}
}

func (cr *cryptocurrencyRepository) GetBalance(symbol string) (float64, error) {
	var balance float64

	err := cr.db.QueryRow("SELECT amount FROM cryptocurrencies WHERE symbol = $1", symbol).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (cr *cryptocurrencyRepository) UpdateBalance(symbol string, amount float64) error {
	_, err := cr.db.Exec("UPDATE cryptocurrencies SET amount = $1 WHERE symbol = $2", amount, symbol)
	if err != nil {
		return err
	}

	return nil
}
