package handlers

import (
	"context"

	"github.com/adelowo/queryapp"
)

func CalculatePrice(ctx context.Context, operation queryapp.Operation, margin float64, exchangeRate float64) (int, error) {
	return 100, nil
}
