package validate_client

import (
	"net/http"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	model "src/server/models"
	//"src/server/response"
	e "src/server/errors"
	"errors"
	c "src/server/constants"
)

func Validate(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var errors []e.Error
		var cl model.Client
		j := r.Header.Get("Client")
		if err1 := json.Unmarshal([]byte(j), &cl); err1 != nil {
			errors = append(errors, e.New("client_error", 1, err1.Error()))
			//response.Failed(w, r, errors, 5)
			return
		}
		_, err2 := govalidator.ValidateStruct(cl)
		if err2 != nil {
			errors = append(errors, e.New("client_error", 2, err2.Error()))
			//response.Failed(w, r, errors, 5)
			return
		}
		if err3 := validateClient(&cl); err3 != nil {
			errors = append(errors, e.New("client_error", 3, err3.Error()))
			//response.Failed(w, r, errors, 5)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func validateClient(client *model.Client) error {
	if  client.Platform != c.IOS && client.Platform != c.Android {
		return errors.New("Invalid validate_client platform. Please update app")
	}
	if client.Version != c.CurrentVersion {
		return errors.New("Invalid validate_client version. Please update app")
	}
	return nil
}

