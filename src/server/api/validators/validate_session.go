package validator

import (
	"gopkg.in/gin-gonic/gin.v1"
	"strings"
	"src/server/response"
	"strconv"
	c "src/server/constants"
	e "src/server/errors"
)

func Session() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var errors []e.Error
		client := ctx.Request.Header.Get("Client")
		values := strings.Split(client, "/")
		if len(values) != 3 {
			errors = append(errors, e.New("header", e.ServerErrorClientHeader, "Invalid 'Client' header"))
			response.Failed(ctx, errors, "Client's header is not exist or format is incorrect")
			ctx.Abort()
			return
		}

		ver := strings.TrimSpace(strings.ToUpper(values[0]))
		platform := strings.TrimSpace(strings.ToUpper(values[1]))
		//_ := strings.TrimSpace(values[2])

		val, err := strconv.ParseFloat(ver, 32)
		if err != nil {
			errors = append(errors, e.New("header", e.ServerErrorClientHeader, "Invalid client version"))
			response.Failed(ctx, errors, "Can't parse version")
			ctx.Abort()

		} else if float32(val) < c.CurrentVersion {
			response.NotSupported(ctx, "Current version bigger than value")
			ctx.Abort()

		} else if platform != c.Android && platform != c.IOS {
			response.NotSupported(ctx, "Invalid a platform's value")
			ctx.Abort()
		}

		// Configure parameters for Context
		ctx.Params = append(ctx.Params, gin.Param{
			Key: kPlatform,
			Value: platform,
		})
		ctx.Params = append(ctx.Params, gin.Param{
			Key: kAppVersion,
			Value: ver,
		})
	}
}
