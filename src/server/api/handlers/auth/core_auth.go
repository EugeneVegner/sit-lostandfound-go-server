package auth

import (
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/response"
	fb "github.com/huandu/facebook"
	//"time"
	"appengine"
	"appengine/datastore"
	e "src/server/errors"
	//model "src/server/models"
	//"github.com/asaskevich/govalidator"
	"log"
	//"go/constant"
	//"time"
	//"src/server/constants"
	//"appengine/urlfetch"
	"appengine/urlfetch"
	//"errors"
	"src/server/models"
)

type inputFB struct {
	FacebookToken string `binding:"required"`
}

type inputEmail struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type facebookUser struct {
	Id        int64 `binding:"required"`
	Name      string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string
	Email     string
}

func FB(c *gin.Context) {

	//c.Abort()
	//var input inputFB
	//var user model.User
	//var fbUser facebookUser
	//var token model.Token
	var errors []e.Error
	log.Println("IS ABORTED: ", c.IsAborted())
	errors = append(errors, e.New("tokenX", 1, "FAK"))
	response.Failed(c, errors)
	//return
	//
	//
	//if err := c.BindJSON(&input); err != nil {
	//	log.Println("BindJSON: ", err.Error())
	//	errors = append(errors, e.New("token", 1, err.Error()))
	//	response.Failed(c, errors)
	//	return
	//}
	//
	//ctx := appengine.NewContext(c.Request)
	//err2 := facebookAuth(ctx, input.FacebookToken, &fbUser)
	//if err2 != nil {
	//	log.Println("facebookAuth: ", err2.Error())
	//	errors = append(errors, e.New("facebook", 2, err2.Error()))
	//	response.Failed(c, errors)
	//	return
	//}
	//
	////err2 := getUserWithToken(ctx, &input, &user, &token)
	////if err2 != nil {
	////	errors = append(errors, e.New("auth_user_error",1, err2.Error()))
	////	response.Failed(c,errors,2)
	////	return
	////}
	//
	//var output struct {
	//	Users  []model.User `json:"users"`
	//	Tokens []model.Token `json:"tokens"`
	//}
	//output.Users = append(output.Users, user)
	//output.Tokens = append(output.Tokens, token)
	//
	//response.Success(c, output)

}

//func getUserWithToken(ctx appengine.Context, input *input, user *model.User, token *model.Token) error {
//
//	//var users []model.User
//
//	users := []model.User{}
//
//	log.Println("input.Email: ",input.Email)
//
//	q := datastore.NewQuery("User").
//	//Ancestor(userKey(ctx)).
//		Filter("Email =",input.Email)
//	//Order("Created")
//	//Filter("Password=",input.Password).
//	//Limit(1)
//
//	_, err := q.GetAll(ctx, &users)
//
//	if err != nil {
//
//		log.Println("getUser: false\n", err.Error(),"\n")
//		return err
//	}
//	//user = nil
//	//log.Println("keys:\n", keys)
//	//log.Println("users:\n", users)
//
//
//	err1 := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
//
//		users := []model.User{}
//
//		log.Println("input.Email: ",input.Email)
//
//		q := datastore.NewQuery("User").
//			Ancestor(userKey(ctx)).
//			Filter("Email =",input.Email)
//		//Order("Created")
//		//Filter("Password=",input.Password).
//		//Limit(1)
//
//		keys, err := q.GetAll(ctx, &users)
//
//		if err != nil {
//
//			log.Println("getUser: false\n", err.Error(),"\n")
//			return err
//		}
//		user = nil
//		log.Println("keys:\n", keys)
//		log.Println("users:\n", users)
//
//		return nil
//
//	}, &datastore.TransactionOptions{
//		XG:       true,
//		Attempts: 3,
//	})
//
//	if err1 != nil {
//		return err1
//	}
//
//	return nil
//
//}

func userKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "User", "", 0, nil)
}

func facebookAuth(ctx appengine.Context, token string, usr *facebookUser) error {

	log.Println("facebookToken: ", token)
	//ses := fb.New(constants.FacebookAppId, constants.FacebookAppSecret)

	//client := urlfetch.Client(ctx)
	//resp, _ := client.Get(uri)
	//doc, _ := goquery.NewDocumentFromResponse(resp)
	//👍 3
	//Sign up for free

	client := urlfetch.Client(ctx)
	fb.Debug = fb.DEBUG_ALL
	fb.SetHttpClient(client)
	fb.Version = "v2.8"

	res, err := fb.Get("/me", fb.Params{
		//"fields": "first_name, email",
		"access_token": token,
	})
	if err != nil {
		return err
	}

	res.DecodeField("id", usr.Id)
	res.DecodeField("name", usr.Name)
	res.DecodeField("first_name", usr.FirstName)
	res.DecodeField("last_name", usr.LastName)
	res.DecodeField("email", usr.Email)
	log.Println("FB Result: ", res)
	log.Println("FB User: ", usr)
	return nil
}

func (fbUser *facebookUser) saveOrUpdateUser(ctx appengine.Context, user *model.User) error {

	users := []model.User{}
	q := datastore.NewQuery("User").Filter("Fib =", fbUser.Id).Limit(1)
	_, err := q.GetAll(ctx, &users)
	if err != nil {
		return err
	}
	if len(users) == 0 {

		//user.SaveUser()

	} else {

		//user = users[0]
		return nil
	}
	return nil

}



//func getUserCredentials(ctx appengine.Context,input *input) (*model.User, *model.Token, *e.Error) {
//
//	usr := new(model.User)
//	tkn := new(model.Token)
//
//	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
//
//		q1 := datastore.NewQuery("User").
//		Ancestor(userKey(ctx)).Filter("Email=",input.Email).Limit(1)
//
//		q2 := datastore.NewQuery("User").
//		Ancestor(userKey(ctx)).Filter("Password=", input.Password).Limit(1)
//
//		var emails []model.User
//		k1, err1 := q1.GetAll(ctx, &emails)
//		if err1 {
//			return err1
//		}
//		log.Println("emails: ",k1)
//
//
//
//
//
//
//		user.EmailVerified = false
//		user.Created = time.Now().UTC().Unix()
//
//		k, err := saveUser(ctx, user)
//		if err != nil {
//			return err
//		}
//
//		_, err2 := saveToken(ctx, token, k)
//		if err2 != nil {
//			return err2
//		}
//
//		return nil
//
//	}, &datastore.TransactionOptions{
//		XG:       true,
//		Attempts: 3,
//	})
//
//	if err != nil {
//		return nil, nil, &e.New("no_user", 5, err.Error())
//	}
//	return usr, tkn, nil
//}

//func getToken(ctx appengine.Context, token *model.Token, user model.User) (*datastore.Key, error) {
//
//	var tokens []model.Token
//	_, err := datastore.NewQuery("Token").
//		Filter("__ID__ =", user.).
//		Filter("Password=",input.Password).
//		Limit(1).
//		GetAll(ctx, users)
//	if err != nil {
//		return nil, err
//	}
//
//
//
//
//
//
//	token.UserId = parentKey.IntID()
//	token.Generate()
//
//
//
//
//	tk := datastore.NewIncompleteKey(ctx, "Token", parentKey)
//	k, err := datastore.Put(ctx, tk, token)
//	if err != nil {
//		return nil, err
//	}
//	token.Id = k.IntID()
//	return k, nil
//
//}

//func getUser(ctx appengine.Context, input *input) (*model.User, error) {
//
//	var users []model.User
//
//	//var usr model.User
//	_, err := datastore.NewQuery("User").
//	//Ancestor(nil).
//	//Filter("Email=",input.Email).
//	//Filter("Password=",input.Password).
//	//Limit(1).
//		GetAll(ctx, &users)
//	//atastore.Get(ctx,keys[0],&usr)
//
//	if err != nil {
//
//		log.Println("getUser: false\n", err.Error())
//		return nil, err
//	}
//	//datastore.Get(ctx,keys[0],&usr)
//
//	if len(users) == 0 {
//		return nil, errors.New("User not found")
//	}
//
//	user := users[0]
//	return &user, nil
//	//
//	//
//	//
//	//t := time.Now().UTC().Unix()
//	//usr.Updated = t
//	//usr.Created = t
//	//usr.EmailVerified = false
//	//
//	//uk := datastore.NewIncompleteKey(ctx, "User", nil)
//	//k, err := datastore.Put(ctx, uk, usr)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//usr.Id = k.IntID()
//	//return k, nil
//}
