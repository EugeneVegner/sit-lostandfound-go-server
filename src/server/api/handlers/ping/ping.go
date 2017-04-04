package ping_handler

import (
	"src/server/router"
	log "src/server/logger"
	"src/server/response"
)

func Test(r *router.Route) {
	log.Func(Test)
	log.Debug("Params: ", r.Context.Params)
	response.Success(r.Context, nil)

}

//var users []model.User
//var lastUpdateTime time.Time
//var mutex sync.RWMutex
//
//func Ping(w http.ResponseWriter, r *http.Request) {
//
//	var m pingObj
//	m.Name = "nameq"
//	m.Title = "dfd"
//	m.Rout = "sdasdasdasdasd"
//
//
//	MakeResponseSuccess(w,r,m)
//
//
//
//
//	//fmt.Fprint(w, "Ping method!\n", r.Method)
//	//
//	//
//	//
//	//
//	//
//	//vars := mux.Vars(r)
//	//todoId := vars["id"]
//	//fmt.Fprintln(w, "\nPing show:", todoId)
//}
//
//func PingToken(w http.ResponseWriter, r *http.Request, user model.User) {
//	//fmt.Fprint(w, "Ping method!\n", r.Method)
//
//	ctx := appengine.NewContext(r)
//	//usr := user.Current(ctx)//  OAuth(ctx,"6")
//
//	if ctx != nil {
//		//fmt.Fprintln(w, "\nPing no user tokrn: ",user.TokenKey)
//		return
//	}
//
//	//appengine.AccessToken()
//
//	//vars := mux.Vars(r)
//	//todoId := vars["id"]
//	fmt.Fprintln(w, "\nPing show: ", user.Id)
//}
