package handlers

import (
	"context"
	"fmt"

	"github.com/adelowo/queryapp"
)

func CalculatePrice(ctx context.Context, btcClient queryapp.Client,
	operation queryapp.Operation, margin float64, exchangeRate float64) (float64, error) {

	if err := queryapp.IsValidOperation(operation); err != nil {
		return 0, err
	}

	switch queryapp.ParseOperation(operation) {
	case queryapp.BUY:
		fmt.Println("Buying")
	case queryapp.SELL:
		fmt.Println("Selling")
	}

	price, err := btcClient.FetchPrice()
	return price, err
}
