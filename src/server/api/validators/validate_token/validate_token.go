package validate_token

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
