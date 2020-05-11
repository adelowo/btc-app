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
	return price, err
}
