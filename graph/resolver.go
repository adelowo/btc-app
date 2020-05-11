package graph

import "github.com/adelowo/queryapp"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	btcClient queryapp.Client
}

func NewResolver(cl queryapp.Client) *Resolver {
	return &Resolver{cl}
}
