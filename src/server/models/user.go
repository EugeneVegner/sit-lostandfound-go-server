package model

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	//"strings"
	"appengine"
	"appengine/datastore"
	"errors"
	"io"
	//"os"
	"encoding/json"
	"time"
	//"go/validate_token"
)

type User struct {
	Id              int64  `json:"id" datastore:"-"`
	Username        string `json:"username" datastore:"," valid:"length(2|32),required"`
	FirstName       string `json:"firstName" datastore:","`
	LastName        string `json:"lastName" datastore:","`
	Name            string `json:"name" datastore:","`
	Email           string `json:"email" datastore:"," valid:"email,required"`
	Password        string `json:"password" datastore:"," valid:"length(8|24),required"`
	EmailVerified   bool   `json:"emailVerified" datastore:","`
	Fid             int64  `json:"fid" datastore:","`
	DeviceToken     string `json:"deviceToken" datastore:","`
	DeviceModel     string `json:"deviceModel" datastore:","`
	PlatformVersion string `json:"platformVersion" datastore:","`
	Platform        string `json:"platform" datastore:","`
	Created         int64  `json:"created"`
	Updated         int64  `json:"updated"`
}

func UserKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "default", 0, nil)
}

func (user *User) key(ctx appengine.Context) *datastore.Key {
	if user.Id == 0 {
		t := time.Now().UTC().Unix()
		user.Updated = t
		user.Created = t
		return datastore.NewIncompleteKey(ctx, "User", nil)
	}
	return datastore.NewKey(ctx, "User", "", user.Id, nil)
}

func (user *User) SaveUser(ctx appengine.Context) (*User, *datastore.Key, error) {
	k, err := datastore.Put(ctx, user.key(ctx), user)
	if err != nil {
		return nil, nil, err
	}
	user.Id = k.IntID()
	user.Updated = time.Now().UTC().Unix()
	return user, k, nil
}

func DecodeUser(r io.ReadCloser) (*User, error) {
	defer r.Close()
	var user User
	err := json.NewDecoder(r).Decode(&user)
	return &user, err
}

func getAll(ctx appengine.Context) ([]User, error) {
	users := []User{}
	ks, err := datastore.NewQuery("User").Order("Created").GetAll(ctx, &users)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		users[i].Id = ks[i].IntID()
	}
	return users, nil
}

//func deleteDoneTodos(c appengine.Context) error {
//	return datastore.RunInTransaction(c, func(c appengine.Context) error {
//		ks, err := datastore.NewQuery("Todo").KeysOnly().Ancestor(defaultUserList(c)).Filter("Done=", true).GetAll(c, nil)
//		if err != nil {
//			return err
//		}
//		return datastore.DeleteMulti(c, ks)
//	}, nil)
//}

//func init() {
//	http.HandleFunc("/todos", handler)
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	c := appengine.NewContext(r)
//	val, err := handleTodos(c, r)
//	if err == nil {
//		err = json.NewEncoder(w).Encode(val)
//	}
//	if err != nil {
//		c.Errorf("todo error: %#v", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//func handleTodos(c appengine.Context, r *http.Request) (interface{}, error) {
//	switch r.Method {
//	case "POST":
//		todo, err := decodeTodo(r.Body)
//		if err != nil {
//			return nil, err
//		}
//		return todo.save(c)
//	case "GET":
//		return getAllTodos(c)
//	case "DELETE":
//		return nil, deleteDoneTodos(c)
//	}
//	return nil, fmt.Errorf("method not implemented")
//}

// RefreshToken refreshes Ttl and Token for the User.
//func (u *User) RefreshToken() error {
//	u.Token.Generate() // = base64.URLEncoding.EncodeToString(validate_token)
//	return nil
//}

//func CurrentUser(w http.ResponseWriter, r *http.Request) (User, utils.Error) {
//
//	s := r.Header.Get("Authorization")
//	if len(s) < 2 {
//		return nil, utils.MakeError("auth", 10, "Not authorized")
//	}
//	ctx := appengine.NewContext(r)
//	tokenKey := datastore.NewKey(ctx, "Token", "key", 0, nil)
//	userKey := datastore.NewKey(ctx, "User", "", 1, tokenKey)
//	var user User
//
//	if err := datastore.Get(ctx, userKey, &user); err != nil {
//		return nil, utils.MakeError("auth", 11, err.Error())
//	}
//
//	return user, nil
//}

//func GetCurrentUser(w http.ResponseWriter, r *http.Request, user interface{}) /**/ utils.Error {
//
//	s := r.Header.Get("Authorization")
//	if len(s) < 2 {
//		return utils.MakeError("auth", 10, "Not authorized")
//	}
//	ctx := appengine.NewContext(r)
//	tokenKey := datastore.NewKey(ctx, "Token", "key", 0, nil)
//	userKey := datastore.NewKey(ctx, "User", "", 1, tokenKey)
//
//	if err := datastore.Get(ctx, userKey, user); err != nil {
//		return utils.MakeError("auth", 11, err.Error())
//	}
//
//	return utils.MakeError("auth", 11, "sdf")
//}

func GetUserByEmail(ctx appengine.Context, email string, user interface{}) error {

	userKey := datastore.NewKey(ctx, "User", "email", 0, nil)

	if err := datastore.Get(ctx, userKey, user); err != nil {
		return err
	}

	return nil
}

func GetUserByEmailUsername(ctx appengine.Context, username string, email string, user interface{}) error {

	q := datastore.NewQuery("User").Filter("Email =", email).Filter("Username =", username)
	c, err := q.Count(ctx)
	if err != nil {
		return errors.New("1")

		//return err
	}

	if c > 0 {
		return errors.New("User allready exist")
	}

	userKey := datastore.NewKey(ctx, "User", "", 0, nil)
	_, errw := datastore.Put(ctx, userKey, user)

	if errw != nil {
		//return errors.New("2")
		return errw
	}

	//var photos []Photo
	//_, err = q.GetAll(ctx, &photos)
	//
	//if err := datastore.Get(ctx, userKey, user); err != nil {
	//	return err
	//}

	return nil
}

//func ValidateToken(h HandlerTokenFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
//		if len(s) != 2 {
//			MakeResponseWithError(w, r, utils.MakeError("", 401, "Not authorized"))
//			return
//		}
//
//		if s[0] != "Basic" {
//			MakeResponseWithError(w, r, utils.MakeError("", 401, "Not authorized"))
//			return
//		}
//
//		tokenKey := s[1]
//
//		ctx := appengine.NewContext(r)
//		usr, err := getUser(ctx, tokenKey)
//		if err != nil {
//			MakeResponseWithError(w, r, utils.MakeError("", 401, err.Error()))
//		}
//
//		h.ServeTokenHTTP(w, r, usr)
//	}
//}

//func CreateUser(w http.ResponseWriter, r *http.Request) {
//
//
//
//	vars := mux.Vars(r)
//	regnr := vars["regnr"]
//	datastore.New
//
//	car := &types.Car{
//		Model:        req.Model,
//		Regnr:        req.Regnr,
//		Year:         req.Year,
//		Type:         req.Type,
//		CreationTime: time.Now(),
//		Sold:         false,
//	}
//	//key := datastore.NewKey(context, "Car", "", 0, nil)
//	carKey := datastore.NewKey(context, "Car", req.Regnr, 0, nil)
//	_, err := datastore.Put(context, carKey, car)
//
//	fmt.Fprint(w, "Welcome!\n")
//}
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func HideUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ping method!\n", r.Method)

	vars := mux.Vars(r)
	todoId := vars["id"]
	fmt.Fprintln(w, "\nPing show:", todoId)
}
