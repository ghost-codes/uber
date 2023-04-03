package middleware

import (

	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(store db.Store) gin.HandlerFunc{
    return func(ctx *gin.Context){
        //but := new(bytes.Buffer)
        //but.ReadFrom(ctx.Request.Body)
       // fmt.Println(but.String())

        ctx.Next()
    }
}
