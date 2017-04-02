package auth

import (
	fb "github.com/huandu/facebook"
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/response"
	//"time"
	"appengine"
	"appengine/datastore"
	e "src/server/errors"
	//model "src/server/models"
	//"github.com/asaskevich/govalidator"
	log "src/server/logger"
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
	Id        string `binding:"required"`
	Name      string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string
	Email     string
}

func FB(c *gin.Context) {
	log.Func(FB)
	var input inputFB
	var errors []e.Error

	if err := c.BindJSON(&input); err != nil {
		log.DebugError("BindJSON: ", err.Error())
		errors = append(errors, e.New("token", 1, err.Error()))
		response.Failed(c, errors, "Can't bind input parameters")
		return
	}

	ctx := appengine.NewContext(c.Request)
	fbUser, err2 := facebookAuth(ctx, input.FacebookToken)
	if err2 != nil {
		log.DebugError("facebookAuth: ", err2.Error())
		errors = append(errors, e.New("facebook", 2, err2.Error()))
		response.Failed(c, errors, "Can't auth to facebook and get user's info")
		return
	}
	log.Debug("fbUser ID: ", fbUser.Id)

	user, token, err3 := configureUserWithToken(ctx, fbUser)
	if err3 != nil {
		errors = append(errors, e.New("facebook", 1, err3.Error()))
		response.Failed(c, errors, "Gan't get/save User and Token")
		return
	}

	var output struct {
		UserId       int64  `json:"userId"`
		Token        string `json:"token"`
		TokenExpired int64  `json:"tokenExpired"`
	}
	output.UserId = user.Id
	output.Token = token.Hash
	output.TokenExpired = token.Expired

	response.Success(c, output)
}

func configureUserWithToken(ctx appengine.Context, fbUser *facebookUser) (*model.User, *model.Token, error) {
	log.Func(configureUserWithToken)
	keys, users, err := model.GetUsersBy(ctx, "Fid=", fbUser.Id, 1)
	if err != nil {
		return nil, nil, err
	}
	if len(users) == 0 {
		return createNewUserAndToken(ctx, fbUser)

	} else {
		user := &users[0]
		userKey := keys[0]
		_, t, err1 := model.GetTokenByUserId(ctx, user.Id)
		if err1 != nil {
			return nil, nil, err1
		}

		t.Generate()
		_, err2 := model.SaveToken(ctx, t, userKey)
		if err2 != nil {
			return nil, nil, err2
		}
		return user, t, nil
	}

	//err1 := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
	//
	//	users := []model.User{}
	//	q := datastore.NewQuery("User").Ancestor(userKey(ctx)).Filter("Fib =",fbUser.Id).Limit(1)
	//	keys, err := q.GetAll(ctx, &users)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(users) == 0 {
	//
	//
	//
	//	} else {
	//
	//
	//	}
	//
	//	if err != nil {
	//
	//		log.Println("getUser: false\n", err.Error(),"\n")
	//		return err
	//	}
	//	user = nil
	//	log.Println("keys:\n", keys)
	//	log.Println("users:\n", users)
	//
	//	return nil
	//
	//}, &datastore.TransactionOptions{
	//	XG:       true,
	//	Attempts: 3,
	//})
	//
	//if err1 != nil {
	//	return err1
	//}

}

func userKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "User", "", 0, nil)
}

func facebookAuth(ctx appengine.Context, token string) (*facebookUser, error) {
	log.Func(facebookAuth)
	log.Debug("facebookToken: ", token)
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
		"fields": "name, first_name, last_name, email",
		"access_token": token,
	})
	if err != nil {
		log.DebugError("Facebook request error: ", err.Error())
		return nil, err
	}
	var user facebookUser
	res.DecodeField("id", &user.Id)
	res.DecodeField("name", &user.Name)
	res.DecodeField("first_name", &user.FirstName)
	res.DecodeField("last_name", &user.LastName)
	res.DecodeField("email", &user.Email)
	log.Debug("FB Result: ", res)
	log.Debug("FB User: ", user)
	return &user, nil
}

func createNewUserAndToken(ctx appengine.Context, fbUser *facebookUser) (*model.User, *model.Token, error) {
	log.Func(createNewUserAndToken)
	var user model.User
	user.Name = fbUser.Name
	user.FirstName = fbUser.FirstName
	user.LastName = fbUser.LastName
	user.Email = fbUser.Email
	if len(user.Email) > 0 {
		user.EmailVerified = true
	}
	user.Fid = fbUser.Id

	var token model.Token
	token.Generate()

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {

		k, err1 := model.SaveUser(ctx, &user)
		if err1 != nil {
			return err1
		}
		token.UserId = user.Id
		_, err2 := model.SaveToken(ctx, &token, k)
		if err2 != nil {
			return err2
		}
		return nil

	}, &datastore.TransactionOptions{
		XG:       true,
		Attempts: 3,
	})

	if err != nil {
		return nil, nil, err
	}
	return &user, &token, nil

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
