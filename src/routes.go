package src

import (
	"net/http"
	//user_client "src/server/api/validators/validate_client"
	//signIn "src/server/api/handlers/sign_in"
	//signUp "src/server/api/handlers/sign_up"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
// ACL
//Route{"SignIn", "POST", "/sign_in", user_client.Validate(signIn.POST)},
//Route{"SignUp", "POST", "/sign_up", user_client.Validate(signUp.POST)},

// User

//Route{"Ping", "GET", "/{id}", handler.BasicAuth(handler.Ping)},
//Route{"PingToken", "GET", "/ping_token", ValidateToken(PingToken)},
//Route{"Ping", "GET", "/ping", Ping},

// User
//Route{"CreateUser", "PUT", "/user", server.models.CreateUser},
}

//var users []model.User
//
//func loadPenguinsJson() {
//	file, err := os.Open("test_users.json")
//	if err != nil {
//		log.Fatal("Can't read users.json:", err)
//		return
//	}
//
//	jsonParser := json.NewDecoder(file)
//	err = jsonParser.Decode(&users)
//	if err != nil {
//		log.Fatal("Can't parse penguins.json:", err)
//		return
//	}
//	// log.Fatal exits the program.
//	// It's important to exit the program if penguins.json can't be read into penguins slice.
//	// Otherwise you will have errors due to unitialised slice.
//}
