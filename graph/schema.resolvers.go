package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/adelowo/queryapp"
	generated1 "github.com/adelowo/queryapp/graph/generated"
	"github.com/adelowo/queryapp/handlers"
)

func (r *queryResolver) CalculatePrice(ctx context.Context, operation queryapp.Operation, margin float64, exchangeRate float64) (float64, error) {
	return handlers.CalculatePrice(ctx, r.btcClient, operation, margin, exchangeRate)
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
