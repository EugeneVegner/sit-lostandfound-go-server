package response

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	e "src/server/errors"
)

type output struct {
	Errors  []e.Error   `json:"errors, omitempty"`
	Data    interface{} `json:"data, omitempty"`
	Rout    string      `json:"rout, omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	send(c, data, nil, http.StatusOK)
}

func Failed(c *gin.Context, errors []e.Error) {
	send(c, nil, errors, http.StatusBadRequest)
}

func ExpiredToken(c *gin.Context) {
	var errors []e.Error
	errors = append(errors, e.New("expired_token",http.StatusForbidden, "Token has been expired"))
	send(c, nil, errors, http.StatusForbidden)
}

func InvalidToken(c *gin.Context) {
	var errors []e.Error
	errors = append(errors, e.New("invalid_token",http.StatusUnauthorized, "Token is invalid"))
	send(c, nil, errors, http.StatusUnauthorized)
}

func NotSupported(c *gin.Context) {
	var errors []e.Error
	errors = append(errors, e.New("unsupport",http.StatusHTTPVersionNotSupported, "Curren app isn't supported"))
	send(c, nil, errors, http.StatusHTTPVersionNotSupported)
}

func send(c *gin.Context, data interface{}, errors []e.Error, code int) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	var resp output
	resp.Data = data
	resp.Errors = errors
	resp.Rout = string(c.Request.URL.Path)

	log.Println("Sent response: ", string(c.Request.URL.Path))
	c.JSON(code, resp)
}



//type HandlerTokenFunc func(http.ResponseWriter, *http.Request, model.User)
//
//// ServeTokenHTTP calls f(w, r).
//func (f HandlerTokenFunc) ServeTokenHTTP(w http.ResponseWriter, r *http.Request, user model.User) {
//	f(w, r, user)
//}
