package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func UserAuthDirective(ctx context.Context, obj interface{},next graphql.Resolver)(interface{},error){
    return nil,nil
}
