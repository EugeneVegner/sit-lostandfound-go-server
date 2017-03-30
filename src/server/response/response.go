package response

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"net/http"
	"encoding/json"
	e "src/server/errors"
	"log"
	model "src/server/models"
)

type handler struct {
	writer http.ResponseWriter
	request *http.Request
}

type output struct {
	Success bool `json:"success"`
	Code    int `json:"code"`
	Errors  []e.Error `json:"errors"`
	Data    interface{} `json:"data"`
	Rout    string `json:"rout,omitempty"`
}

//func (output *output) AddError(err e.Error ) []e.Error   {
//	response.Errors = append(response.Errors, err)
//	return response.Errors
//}

func Success(c *gin.Context, data interface{}) {

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	var resp output
	resp.Success = true
	resp.Code = 0 // NoErrors
	resp.Errors = nil
	resp.Data = data
	resp.Rout = string(c.Request.URL.Path)

	log.Println("MakeResponseSuccess: ")

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println("MakeResponseSuccess: false\n", string(body))
		fmt.Fprint(c.Writer, string("ADslkajsdlk"))
		return
	}
	log.Println("MakeResponseSuccess: true\n", string(body))
	//fmt.Fprint(c.Writer, string(body))

	c.JSON(http.StatusOK, body)

}

//func Failed(w http.ResponseWriter, r *http.Request, err e.Error ) {
//	log.Println("MakeResponseError: ", err.Message)
//	errs := []e.Error {err}
//	FailedWithErrors(w, r, errs, err.Code)
//}

func Failed(c *gin.Context, errors []e.Error, code int) {

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	var resp output
	resp.Success = false
	resp.Code = code // Error code
	resp.Data = nil
	resp.Rout = string(c.Request.URL.Path)

	_, err := json.Marshal(errors)
	if err != nil || errors == nil {
		resp.Errors = append(resp.Errors, e.Unknown(code))
	} else {
		resp.Errors = errors
	}

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println("MakeResponseError: true\n", body)
	}
	log.Println("MakeResponseError: true\n", body)
	//	log.Println("MakeResponseError: false\n", string(body))
	//	//fmt.Fprint(w, string(body))
	//	resp.Errors = append(resp.Errors, e.Unknown(code))
	//	c.JSON(http.StatusOK, &body)
	//	return
	//}
	log.Println("MakeResponseError: true\n", )
	//fmt.Fprint(w, string(body))
	c.JSON(http.StatusOK, resp)

}

type HandlerTokenFunc func(http.ResponseWriter, *http.Request, model.User)

// ServeTokenHTTP calls f(w, r).
func (f HandlerTokenFunc) ServeTokenHTTP(w http.ResponseWriter, r *http.Request, user model.User) {
	f(w, r, user)
}
