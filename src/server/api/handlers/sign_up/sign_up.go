package sign_up_handler

import (
	"time"
	//"net/http"
	"appengine"
	"appengine/datastore"
	"gopkg.in/gin-gonic/gin.v1"
	e "src/server/errors"
	"src/server/models"
	"src/server/response"
	"src/server/utils"
)

type output struct {
	Users  []model.User  `json:"users, required" binding:"required"`
	Tokens []model.Token `json:"tokens, required" binding:"required"`
}

func POST(c *gin.Context) {
	var user model.User
	var token model.Token
	errors := utils.EncodeBody(c, &user)
	if errors != nil {
		response.Failed(c, errors, "The body can't be encoded")
		return
	}

	ctx := appengine.NewContext(c.Request)
	err := createNewUserWithToken(ctx, &user, &token)
	if err != nil {
		response.Failed(c, err, "Can't create new User with Token")
		return
	}

	var output output
	output.Users = append(output.Users, user)
	output.Tokens = append(output.Tokens, token)

	response.Success(c, output)
}

func createNewUserWithToken(ctx appengine.Context, user *model.User, token *model.Token) []e.Error {

	var users []model.User
	var errors []e.Error
	q := datastore.NewQuery("User")
	_, err2 := q.Filter("Email=", user.Email).Limit(1).GetAll(ctx, &users)
	if err2 != nil {
		errors = append(errors, e.New("email", 33, err2.Error()))
		return errors
	}
	if len(users) > 0 {
		errors = append(errors, e.New("email", 34, "Email already exist"))
		return errors
	}

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {

		user.EmailVerified = false
		user.Created = time.Now().UTC().Unix()

		k, err := saveUser(ctx, user)
		if err != nil {
			return err
		}

		_, err2 := saveToken(ctx, token, k)
		if err2 != nil {
			return err2
		}

		return nil

	}, &datastore.TransactionOptions{
		XG:       true,
		Attempts: 3,
	})

	if err != nil {
		errors = append(errors, e.New("registration", 35, err.Error()))
		return errors
	}

	return nil

}

func saveToken(ctx appengine.Context, token *model.Token, parentKey *datastore.Key) (*datastore.Key, error) {

	token.UserId = parentKey.IntID()
	token.Generate()
	tk := datastore.NewIncompleteKey(ctx, "Token", parentKey)
	k, err := datastore.Put(ctx, tk, token)
	if err != nil {
		return nil, err
	}
	token.Id = k.IntID()
	return k, nil

}

func saveUser(ctx appengine.Context, usr *model.User) (*datastore.Key, error) {

	t := time.Now().UTC().Unix()
	usr.Updated = t
	usr.Created = t
	usr.EmailVerified = false

	uk := datastore.NewIncompleteKey(ctx, "User", nil)
	k, err := datastore.Put(ctx, uk, usr)
	if err != nil {
		return nil, err
	}
	usr.Id = k.IntID()
	return k, nil
}
