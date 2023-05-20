package dataloader

import (
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/gin-gonic/gin"
)

func DataloaderMiddleware(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loaders := newLoader(ctx, store)
		// augmentedCtx := context.WithValue(ctx, key, loaders)
		ctx.Set(key, loaders)

		ctx.Next()
	}
}
