package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ghost-codes/uber/middleware"
	"github.com/ghost-codes/uber/util"
)

func UserAuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	ginCtx, err := util.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

    _,ok:= ginCtx.Get(middleware.UserPayloadKey)
    if !ok{
        return nil, fmt.Errorf("Access denied\nUnauthenticated")
    }

    
	userType, ok := ginCtx.Get(middleware.UserTypeKey)

	if !ok || userType != middleware.UserTypeClient {
		return nil, fmt.Errorf("Access denied: user must be of type client")
	}

	return next(ctx)
}
