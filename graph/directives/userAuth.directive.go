package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ghost-codes/uber/graph/model"
)



func UserAuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, typeArg model.Type) (res interface{}, err error){
    return nil,nil
}
