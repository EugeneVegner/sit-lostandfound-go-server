package fb_auth

import (
	"src/server/response"
	"gopkg.in/gin-gonic/gin.v1"
	fb "github.com/huandu/facebook"
	model "src/server/models"
	"appengine"
	"log"
	e "src/server/errors"
	"src/server/constants"

)

func GET(c *gin.Context) {

	var user model.User
	var token model.Token

	ctx := appengine.NewContext(c.Request)


	errs := facebookAuth(ctx, c.Param("token"))
	if errs {
		response.Failed(c, errs,30)
		return
	}

	var output struct {
		Users  []model.User  `json:"users"`
		Tokens []model.Token `json:"tokens"`
	}
	output.Users = append(output.Users, user)
	output.Tokens = append(output.Tokens, token)

	response.Success(c, output)


}

func createUserIfNeeded(ctx appengine.Context, user interface{}) error {


	return nil
}


func facebookAuth(ctx appengine.Context, token string) []e.Error {

	var errors []e.Error

	if len(token) == 0 {
		errors = append(errors, e.New("fb", 31, "Invalid facebook token"))
		return errors
	}

	fb.Debug = fb.DEBUG_ALL
	ses := fb.New(constants.FacebookAppId, constants.FacebookAppSecret)

	ses.re

	res, _ := fb.Get("/538744468", fb.Params{
		"fields": "first_name",
		"access_token": token,
	})

	//
//appId      : '319017188429459',
//xfbml      : true,
//	version    : 'v2.8'


	var first_name string
	var last_name string
	var name string
	var email string

	res.DecodeField("name", &name)
	res.DecodeField("first_name", &first_name)
	res.DecodeField("last_name", &last_name)
	res.DecodeField("email", &email)

	log.Println("input.Email: ", res)

	return nil
}


