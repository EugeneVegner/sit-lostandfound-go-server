package api


//type handler struct {
//	writer http.ResponseWriter
//	request *http.Request
//}
//
//type response struct {
//	Success bool `json:"success"`
//	Code    int `json:"code"`
//	Errors  []e.Error `json:"errors"`
//	Data    interface{} `json:"data"`
//	Rout    string `json:"rout,omitempty"`
//}
//
//func (response *response) AddError(err e.Error ) []e.Error   {
//	response.Errors = append(response.Errors, err)
//	return response.Errors
//}
//
//func MakeResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("X-Content-Type-Options", "nosniff")
//
//	var resp response
//	resp.Success = true
//	resp.Code = 0 // NoErrors
//	resp.Errors = nil
//	resp.Data = data
//	resp.Rout = string(r.URL.Path)
//
//	log.Println("MakeResponseSuccess: ")
//
//	body, err := json.Marshal(resp)
//	if err != nil {
//		log.Println("MakeResponseSuccess: false\n", string(body))
//		fmt.Fprint(w, string("ADslkajsdlk"))
//		return
//	}
//	log.Println("MakeResponseSuccess: true\n", string(body))
//	fmt.Fprint(w, string(body))
//
//
//}
//
//func MakeResponseWithError(w http.ResponseWriter, r *http.Request, err e.Error ) {
//	log.Println("MakeResponseError: ", err.Message)
//	errs := []e.Error {err}
//	MakeResponseWithErrors(w, r, errs, err.Code)
//}
//
//func MakeResponseWithErrors(w http.ResponseWriter, r *http.Request, errors []e.Error, code int) {
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Header().Set("X-Content-Type-Options", "nosniff")
//
//	var resp response
//	resp.Success = false
//	resp.Code = code // Error code
//	resp.Data = nil
//	resp.Rout = string(r.URL.Path)
//
//	_, err := json.Marshal(errors)
//	if err != nil || errors == nil {
//		resp.Errors = append(resp.Errors, *e.UnknownError(code))
//	} else {
//		resp.Errors = errors
//	}
//
//	body, err := json.Marshal(resp)
//	if err != nil {
//		log.Println("MakeResponseError: false\n", string(body))
//		fmt.Fprint(w, string(body))
//		return
//	}
//	log.Println("MakeResponseError: true\n", string(body))
//	fmt.Fprint(w, string(body))
//
//}
//
//type HandlerTokenFunc func(http.ResponseWriter, *http.Request, model.User)
//
//// ServeTokenHTTP calls f(w, r).
//func (f HandlerTokenFunc) ServeTokenHTTP(w http.ResponseWriter, r *http.Request, user model.User) {
//	f(w, r, user)
//}

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
//
//
//func ClientHeader(h http.HandlerFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		client := model.Client{}
//		ag := r.Header.Get("Client")
//		if parser_error := json.Unmarshal([]byte(ag), &client); parser_error != nil {
//			MakeResponseWithError(w, r, *e.MakeError("client_error", 1, parser_error.Error()))
//			return
//		}
//		_, err1 := govalidator.ValidateStruct(client)
//		if err1 != nil {
//			MakeResponseWithError(w, r, *e.MakeError("client_error", 2, err1.Error()))
//			return
//		}
//		if err := validateClient(&client); err != nil {
//			MakeResponseWithError(w, r, *e.MakeError("client_error", 3, err.Error()))
//			return
//		}
//		h.ServeHTTP(w, r)
//	}
//}
//
//func validateClient(client *model.Client) error {
//	if  client.Platform != c.IOS && client.Platform != c.Android {
//		return errors.New("Invalid validate_client platform. Please update app")
//	}
//	if client.Version != c.CurrentVersion {
//		return errors.New("Invalid validate_client version. Please update app")
//	}
//	return nil
//}
