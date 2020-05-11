package handlers

import (
	"context"

	"github.com/adelowo/queryapp"
)

func CalculatePrice(ctx context.Context, btcClient queryapp.Client,
	operation queryapp.Operation, margin float64, exchangeRate float64) (float64, error) {

	if err := queryapp.IsValidOperation(operation); err != nil {
		return 0, err
	}

	price, err := btcClient.FetchPrice()
	if err != nil {
		return 0, err
	}

	var amount float64

	switch queryapp.ParseOperation(operation) {
	case queryapp.BUY:
		amount = calculate(price, margin, true)
	case queryapp.SELL:
		amount = calculate(price, margin, false)
	}

	return amount * exchangeRate, nil
}

// feature flags like plusOp are kind of terrible but :)
func calculate(btcPrice, margin float64, plusOp bool) float64 {
	x := (margin / 100) * btcPrice

	if plusOp {
		return btcPrice + x
	}

	return btcPrice - x
}
