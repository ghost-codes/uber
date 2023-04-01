package graph

import (
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/util"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    Store *db.Store
    Config util.Config
}
