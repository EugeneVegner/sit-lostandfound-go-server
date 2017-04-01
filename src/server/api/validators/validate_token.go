package validator

import (
	"gopkg.in/gin-gonic/gin.v1"
)

//import "net/http"

//func ValidateToken(h HandlerTokenFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		//s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
//		//if len(s) != 2 {
//		//	MakeResponseWithError(w,r, utils.MakeError("", 401, "Not authorized"))
//		//	return
//		//}
//		//
//		//if s[0] != "Basic" {
//		//	MakeResponseWithError(w,r, utils.MakeError("", 401, "Not authorized"))
//		//	return
//		//}
//		//
//		//tokenKey := s[1]
//		//
//		//ctx := appengine.NewContext(r)
//		//
//		//
//		// usr, err := model.CurrentUser(w,r)
//		//if err != nil {
//		//	MakeResponseWithError(w,r, utils.MakeError("", 12, err))
//		//}
//
//		//usr, err := model.CurrentUser(w,r)
//		//h.ServeTokenHTTP(w, r, nil)
//	}
//}


func Token() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Request.Header


		//var errors []e.Error
		//var cl model.Client
		//j := r.Header.Get("Client")
		//if err1 := json.Unmarshal([]byte(j), &cl); err1 != nil {
		//	errors = append(errors, e.New("client_error", 1, err1.Error()))
		//	//response.Failed(w, r, errors, 5)
		//	return
		//}
		//_, err2 := govalidator.ValidateStruct(cl)
		//if err2 != nil {
		//	errors = append(errors, e.New("client_error", 2, err2.Error()))
		//	//response.Failed(w, r, errors, 5)
		//	return
		//}
		//if err3 := validateClient(&cl); err3 != nil {
		//	errors = append(errors, e.New("client_error", 3, err3.Error()))
		//	//response.Failed(w, r, errors, 5)
		//	return
		//}
		//
		//
		//c.Next()
	}
}