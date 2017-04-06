package auth

//EAAGdJ4AWqAQBAN3WSXD6EGdZCO8PphlsEf7iairkhrUdCnTrN3347ojiMZCIYKv7InTx5nPWd0mmfGHBMvd3yDBPfi4ZBAAV6O51ZAKBclYUV6JICnDgdMzzZCfU5HoB1xSOjWuGMaZBSmIEkMbXs2uf2wqFcVZCfok8NBkAg6yBZAsDQohfiGhf8pxxx6Fe1C69fD9xEMD9o8bbO8vhf6CeKNiLPbs6yLsZD

import (
	fb "github.com/huandu/facebook"
	"gopkg.in/gin-gonic/gin.v1"
	"src/server/response"
	"appengine"
	"appengine/datastore"
	c "src/server/constants"
	e "src/server/errors"
	log "src/server/logger"
	"appengine/urlfetch"
	"src/server/models"
)

type input struct {
	FacebookToken string `binding:"required"`
	DeviceToken string
	DeviceID string
}

type output struct {
	UserId int64 `binding:"required"`
	Token string `binding:"required"`
	Expired string `binding:"required"`
}

type facebookUser struct {
	Id        string `binding:"required"`
	Name      string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string
	Email     string
}

func FB(ctx *gin.Context) {
	log.Func(FB)
	var input input
	var errors []e.Error

	if err := ctx.BindJSON(&input); err != nil {
		log.DebugError("BindJSON: ", err.Error())
		errors = append(errors, e.New("token", 1, err.Error()))
		response.Failed(ctx, errors, "Can't bind input parameters")
		return
	}

	db := appengine.NewContext(ctx.Request)
	fbUser, err2 := facebookAuth(db, input.FacebookToken)
	if err2 != nil {
		log.DebugError("facebookAuth: ", err2.Error())
		errors = append(errors, e.New("facebook", 2, err2.Error()))
		response.Failed(ctx, errors, "Can't auth to facebook and get user's info")
		return
	}
	log.Debug("fbUser ID: ", fbUser.Id)

	user, session, err3 := updateUserAtSession(db, ctx, fbUser)
	if err3 != nil {
		errors = append(errors, e.New("facebook", 1, err3.Error()))
		response.Failed(ctx, errors, "Gan't get/save User and Token")
		return
	}

	var output struct {
		UserId       int64  `json:"userId"`
		Token        string `json:"token"`
		TokenExpired int64  `json:"tokenExpired"`
	}
	output.UserId = user.Id
	output.Token = session.Token.Hash
	output.TokenExpired = session.Token.Expired

	response.Success(ctx, output)
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

func updateUserAtSession(db appengine.Context, ctx *gin.Context, fbUser *facebookUser) (*model.User, *model.Session, error) {
	log.Func(updateUserAtSession)
	keys, users, err := model.GetUsersBy(db, "Fid=", fbUser.Id, 1)
	if err != nil {
		return nil, nil, err
	}
	if len(users) == 0 {
		return createNewUserWithNewSession(db, ctx, fbUser)

	} else {
		user := &users[0]
		userKey := keys[0]
		_, s, err1 := model.GetSessionByUserId(db, user.Id)
		if err1 != nil {
			return nil, nil, err1
		}

		s.Token.Generate()
		_, err2 := model.SaveSession(db, s, userKey)
		if err2 != nil {
			return nil, nil, err2
		}
		return user, s, nil
	}

}

func createNewUserWithNewSession(db appengine.Context, ctx *gin.Context, fbUser *facebookUser) (*model.User, *model.Session, error) {
	log.Func(createNewUserWithNewSession)
	var user model.User
	user.Name = fbUser.Name
	user.FirstName = fbUser.FirstName
	user.LastName = fbUser.LastName
	user.Email = fbUser.Email
	user.Provider = c.ProviderTypeFB
	if len(user.Email) > 0 {
		user.EmailVerified = true
	}
	user.Fid = fbUser.Id

	log.Debug("Try to create session")
	var session model.Session
	var device model.Device
	if err := device.New(ctx); err != nil {
		return nil, nil, err
	}
	var token model.Token
	if err := token.New(ctx); err != nil {
		return nil, nil, err
	}
	token.Generate()

	session.Device = device
	session.Token = token
	log.Debug("Session has been created")

	err := datastore.RunInTransaction(db, func(ctx appengine.Context) error {
		log.Debug("RunInTransaction")
		k, err1 := model.SaveUser(ctx, &user)
		if err1 != nil {
			return err1
		}
		session.UserId = user.Id
		_, err2 := model.SaveSession(ctx, &session, k)
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
	return &user, &session, nil

}