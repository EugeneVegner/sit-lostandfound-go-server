package validator

import (
	"appengine"
	"appengine/datastore"
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/models"
	"src/server/response"
	"strings"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {

		token_key := c.Request.Header.Get("Authorization")
		if len(strings.TrimSpace(token_key)) == 0 {
			response.InvalidToken(c, "'Authorization' header not exist or Token is not exist")
			c.Abort()
			return
		}

		var tokens []model.Token
		ctx := appengine.NewContext(c.Request)
		q := datastore.NewQuery("Token").Filter("Key =", token_key).Limit(1)

		_, err := q.GetAll(ctx, &tokens)
		if err != nil {
			response.InvalidToken(c, "Token query error: "+err.Error())
			c.Abort()
			return
		}
		if len(tokens) == 0 {
			response.InvalidToken(c, "Token not found")
			c.Abort()
			return
		}
		token := tokens[0]
		if token.IsExpired() {
			response.ExpiredToken(c, "Token is expired")
			c.Abort()
			return
		}

	}
}
