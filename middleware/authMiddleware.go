package middleware

import (
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
    authorizationKey="Authorization"
    authorizationBearerType="bearer"
    userPayloadKey="client";
)
 
func AuthMiddleware(store db.Store,firebaseAuth *auth.Client) gin.HandlerFunc{
    return func(ctx *gin.Context){
        authorizationHeader := ctx.GetHeader(authorizationKey);
        if len(authorizationHeader)==0{
            ctx.Next();
            return ;
        }
        fields:= strings.Fields(authorizationHeader)
        
        if len(fields)<2{
            ctx.Next();
            return;
        }

        if strings.ToLower(fields[0]) != authorizationBearerType{
            ctx.Next();
            return;
        }

        token,err:=firebaseAuth.VerifyIDToken(ctx,fields[1]);
        if err!=nil{
            ctx.AbortWithError(http.StatusForbidden,err);
            return;
        }
        user,err:=store.FetchUserMetaDataByID(ctx,token.UID);
        if err!=nil{
            ctx.AbortWithError(http.StatusForbidden,err);
            return;
        }

        ctx.Set(userPayloadKey,user)

        ctx.Next()
    }
}
