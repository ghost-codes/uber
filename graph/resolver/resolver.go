package graph

import (
	"firebase.google.com/go/auth"
	"github.com/ghost-codes/uber/dataloader"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/token"
	"github.com/ghost-codes/uber/util"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	FirebaseAuth *auth.Client
	Store        *db.Store
	Maker        token.Maker
	Config       util.Config
	DataLoaders  dataloader.Retriever
}
