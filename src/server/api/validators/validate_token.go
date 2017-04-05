package validator

import (
	"appengine"
	"appengine/datastore"
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/models"
	"src/server/response"
	"strings"
	c "src/server/constants"
)


func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token_key := ctx.Request.Header.Get("Authorization")
		if len(strings.TrimSpace(token_key)) == 0 {
			response.InvalidToken(ctx, "'Authorization' header not exist or Token is not exist")
			ctx.Abort()
			return
		}

		var tokens []model.Token
		ctx_req := appengine.NewContext(ctx.Request)
		q := datastore.NewQuery("Token").Filter("Key =", token_key).Limit(1)

		_, err := q.GetAll(ctx_req, &tokens)
		if err != nil {
			response.InvalidToken(ctx, "Token query error: "+err.Error())
			ctx.Abort()
			return
		}
		if len(tokens) == 0 {
			response.InvalidToken(ctx, "Token not found")
			ctx.Abort()
			return
		}
		token := tokens[0]
		if token.IsExpired() {
			response.ExpiredToken(ctx, "Token is expired")
			ctx.Abort()
			return
		}

		// Configure parameters for Context
		ctx.Params = append(ctx.Params, gin.Param{
			Key: c.ParamKeyUserId,
			Value: string(token.UserId),

		})
		ctx.Params = append(ctx.Params, gin.Param{
			Key: c.ParamKeySessionToken,
			Value: string(token.Hash),
		})


	}
}
