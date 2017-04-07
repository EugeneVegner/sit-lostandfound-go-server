package auth

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
	"src/server/api"
	"src/server/utils"
)

type input struct {
	FacebookToken string `binding:"required"`
	DeviceId      string `binding:"required"`
	DeviceToken   string
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
	errors := utils.EncodeBody(ctx, &input)
	if errors != nil {
		response.Failed(ctx, errors, "The body can't be encoded")
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

	log.Debug("Try to find user by fid: ", fbUser.Id)
	userKey, user, err1 := model.GetUserBy(db, "Fid=", fbUser.Id)
	if err1 != nil {
		log.Debug("err1: ", err1.Error())
		errors = append(errors, e.New("user", e.ServerErrorEmailNotFound, err1.Error()))
		response.Failed(ctx, errors, err1.Error())
		return
	}

	var session *model.Session
	if user == nil {
		log.Debug("err1: user not found")
		u, s, err2 := createNewUserWithNewSession(db, ctx, fbUser, input.DeviceId, input.DeviceToken)
		if err2 != nil {
			log.Debug("err2: ", err2.Error())
			//errors = append(errors, e.New("err2", e.ServerErrorNewEntity, err2.Error()))
			response.Failed(ctx, utils.ReflectError(err2), err2.Error())
			return
		}
		user = u
		session = s

	} else {
		_, s, err3 := api.GetAndUpdateSessionIfNeeded(db, userKey, input.DeviceId, input.DeviceToken)
		if err3 != nil {
			log.Debug("err3: ", err3.Error())
			//errors = utils.ReflectError(err3)
			response.Failed(ctx, utils.ReflectError(err3), err3.Error())
			return
		}
		session = s
	}

	var output api.AuthOutput
	output.UserId = user.Id
	output.Token = session.Token
	output.Expired = session.Expired
	response.Success(ctx, output)
}

func facebookAuth(ctx appengine.Context, token string) (*facebookUser, error) {
	log.Func(facebookAuth)
	log.Debug("facebookToken: ", token)

	client := urlfetch.Client(ctx)
	fb.Debug = fb.DEBUG_ALL
	fb.SetHttpClient(client)
	fb.Version = "v2.8"

	res, err := fb.Get("/me", fb.Params{
		"fields":       "name, first_name, last_name, email",
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

func createNewUserWithNewSession(db appengine.Context, ctx *gin.Context, fbUser *facebookUser, deviceId string, deviceToken string) (*model.User, *model.Session, error) {
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
	err := datastore.RunInTransaction(db, func(ctx appengine.Context) error {
		log.Debug("RunInTransaction")
		userKey, err1 := model.SaveUser(ctx, &user)
		if err1 != nil {
			return err1
		}
		_, s, err2 := api.GetAndUpdateSessionIfNeeded(db, userKey, deviceId, deviceToken)
		if err2 != nil {
			log.Debug("err2: ", err2.Error())
			return err2
		}
		session = *s
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
