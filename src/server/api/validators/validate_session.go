package validator

import (
	"gopkg.in/gin-gonic/gin.v1"
	"strings"
	"src/server/response"
	//"strconv"
	c "src/server/constants"
	//e "src/server/errors"
	"src/server/models"
	"appengine"
	"appengine/datastore"
)

func Session() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token_key := ctx.Request.Header.Get("Authorization")
		if len(strings.TrimSpace(token_key)) == 0 {
			response.InvalidToken(ctx, "'Authorization' header not exist or Session is not exist")
			ctx.Abort()
			return
		}

		var sessions []model.Session
		ctx_req := appengine.NewContext(ctx.Request)
		q := datastore.NewQuery("Session").Filter("Token =", token_key).Limit(1)

		_, err := q.GetAll(ctx_req, &sessions)
		if err != nil {
			response.InvalidToken(ctx, "Session query error: "+err.Error())
			ctx.Abort()
			return
		}
		if len(sessions) == 0 {
			response.InvalidToken(ctx, "Session not found")
			ctx.Abort()
			return
		}
		session := sessions[0]
		if session.Token.IsExpired() {
			response.ExpiredToken(ctx, "Token is expired")
			ctx.Abort()
			return
		}

		// Configure parameters for Context
		ctx.Params = append(ctx.Params, gin.Param{
			Key: c.ParamKeyUserId,
			Value: string(session.UserId),

		})
		ctx.Params = append(ctx.Params, gin.Param{
			Key: c.ParamKeySessionToken,
			Value: session.Token.Hash,
		})

	}
}
