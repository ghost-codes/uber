package dataloader

import (
	"context"

	db "github.com/ghost-codes/uber/db/sqlc"
)

// type  string

const key = "dataloaders"

// Loaders holds references to individual dataloaders
type Loaders struct {
}

// initialize loaders
func newLoader(ctx context.Context, store db.Store) *Loaders {
	return &Loaders{}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key string
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}
