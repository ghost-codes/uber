package directives

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ghost-codes/uber/token"
	"github.com/ghost-codes/uber/util"
)

const (
	authorizationKey        = "Authorization"
	authorizationBearerType = "bearer"
	DriverPayloadKey        = "driver"
)

type DirectiveAuth func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

func AuthenticateDriverfunc(maker token.Maker) DirectiveAuth {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ginCtx, err := util.GinContextFromContext(ctx)
		if err != nil {
			return nil, err
		}

		authorizationHeader := ginCtx.GetHeader(authorizationKey)
		if len(authorizationHeader) == 0 {
			return nil, fmt.Errorf("Access denied\nUnauthenticated")
		}
		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			return nil, fmt.Errorf("Access denied\nUnauthenticated")
		}

		if strings.ToLower(fields[0]) != authorizationBearerType {
			return nil, fmt.Errorf("Access denied\nUnauthenticated")
		}
		token := fields[1]

		payload, err := maker.VerifyToken(token)
		if err != nil {
			return nil, err
		}

		ginCtx.Set(DriverPayloadKey, payload)
		return next(ctx)
	}
}
