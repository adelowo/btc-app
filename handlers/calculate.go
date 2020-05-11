package handlers

import (
	"context"

	"github.com/adelowo/queryapp"
)

func CalculatePrice(ctx context.Context, btcClient queryapp.Client,
	operation queryapp.Operation, margin float64, exchangeRate float64) (float64, error) {

	price, err := btcClient.FetchPrice()
	return price, err
}
