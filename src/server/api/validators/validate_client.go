package validator

import (
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/constants"
	e "src/server/errors"
	"src/server/response"
	"strconv"
	"strings"
)

func Client() gin.HandlerFunc {
	return func(c *gin.Context) {

		var errors []e.Error
		client := c.Request.Header.Get("Client")
		values := strings.Split(client, "/")
		if len(values) != 3 {
			errors = append(errors, e.New("header", e.ServerErrorClientHeader, "Invalid 'Client' header"))
			response.Failed(c, errors, "Client's header is not exist or format is incorrect")
			c.Abort()
			return
		}

		ver := strings.TrimSpace(strings.ToUpper(values[0]))
		platform := strings.TrimSpace(strings.ToUpper(values[1]))
		//_ := strings.TrimSpace(values[2])

		val, err := strconv.ParseFloat(ver, 32)
		if err != nil {
			errors = append(errors, e.New("header", e.ServerErrorClientHeader, "Invalid client version"))
			response.Failed(c, errors, "Can't parse version")
			c.Abort()

		} else if float32(val) < constants.CurrentVersion {
			response.NotSupported(c, "Current version bigger than value")
			c.Abort()

		} else if platform != constants.Android && platform != constants.IOS {
			response.NotSupported(c, "Invalid a platform's value")
			c.Abort()
		}
	}
}
