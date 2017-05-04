package sign_up_handler

import (
	"gopkg.in/gin-gonic/gin.v1"
	"appengine"
	"appengine/datastore"
	e "src/server/errors"
	c "src/server/constants"
	"src/server/models"
	"src/server/response"
	"src/server/api"
	log "src/server/logger"
	"src/server/utils"
	"github.com/asaskevich/govalidator"
)

type input struct {
	FirstName   string `valid:"required" binding:"required"`
	SecondName  string `valid:"required" binding:"required"`
	Email       string `valid:"required" binding:"required"`
	Password    string `valid:"required" binding:"required"`
	DeviceId    string `valid:"required" binding:"required"`
	DeviceToken string
}

var users []model.User

func POST(ctx *gin.Context) {
	log.Func(POST)

	var input input
	errors := utils.EncodeBody(ctx, input)
	if errors != nil {
		response.Failed(ctx, errors, "The body can't be encoded")
		return
	}

	db := appengine.NewContext(ctx.Request)

	// Check user on exist
	q := datastore.NewQuery("User")
	_, err1 := q.Filter("Email=", input.Email).Limit(1).GetAll(db, &users)
	if err1 != nil {
		errors = append(errors, e.New("email", e.ServerErrorEmailExist, err1.Error()))
		log.Debug("User with email <%v> already exist", input.Email)
		response.Failed(ctx, errors, "This email already exist")
		return
	}

	var user model.User
	user.Email = input.Email
	user.EmailVerified = false
	user.FirstName = input.FirstName
	user.LastName = input.SecondName
	user.Name = user.FirstName + " " + user.LastName
	user.Password = input.Password

	if _, err := govalidator.ValidateStruct(user); err != nil {
		response.Failed(ctx, utils.ReflectError(err), "User object is invalid")
		return
	}

	var session *model.Session

	err := datastore.RunInTransaction(db, func(db appengine.Context) error {
		log.Debug("RunInTransaction")

		userKey, err1 := model.SaveUser(db, &user)
		if err1 != nil {
			log.Debug("err1: ",err1.Error())
			return err1
		}

		_, s, err2 := api.GetAndUpdateSessionIfNeeded(
			db,
			userKey,
			input.DeviceId,
			input.DeviceToken,
			ctx.Param(c.ParamKeyClientPlatform))

		if err2 != nil {
			log.Debug("err2: ",err2.Error())
			return err2
		}
		session = s
		return nil

	}, &datastore.TransactionOptions{
		XG:       true,
		Attempts: 3,
	})

	if err != nil {
		errors = append(errors, e.New("save_entity", e.ServerErrorNewEntity, "Can't create new User with Session"))
		response.Failed(ctx, errors, "Can't create new User with Token")
		return
	}

	var output api.AuthOutput
	output.UserId = user.Id
	output.Token = session.Token
	output.Expired = session.Expired
	response.Success(ctx, output)
}