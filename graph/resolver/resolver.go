package graph

import (
	"firebase.google.com/go/auth"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/util"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    FirebaseAuth *auth.Client
    Store *db.Store
    Config util.Config

}
