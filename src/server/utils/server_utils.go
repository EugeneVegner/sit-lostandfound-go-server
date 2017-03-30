package utils

import (
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
	e "src/server/errors"
	//"io/ioutil"
	//"io"
	//"encoding/json"
	"github.com/asaskevich/govalidator"
	"strings"
	"log"
)

func EncodeBody(c *gin.Context, input interface{}) []e.Error {
	var errors []e.Error
	//body, err := ioutil.ReadAll(r.Body)
	//log.Println("r.Body: ", r.Body)
	//log.Println("body: ", body)
	//log.Println("body err: ", err)

	if err := c.BindJSON(&input); err != nil {
		errors = append(errors, e.New("body_error",1, err.Error()))
		return errors
	}




	//switch {
	//case err == io.EOF:
	//	log.Println("body_error[1]: EOF")
	//	errors = append(errors, e.New("body_error",1, "Request body is empty. EOF"))
	//	return errors
	//
	//case err != nil:
	//	log.Println("body_error[2]: ", err)
	//	errors = append(errors, e.New("body_error",2, err.Error()))
	//	return errors
	//}
	//
	////err := json.NewDecoder(r.Body).Decode(input)
	//if parser_error := json.Unmarshal(body, input); parser_error != nil {
	//	log.Println("body_error[3]: ", parser_error)
	//	errors = append(errors, e.New("body_error", 3, parser_error.Error()))
	//	return errors
	//}

	// Validate input structure
	_, struct_error := govalidator.ValidateStruct(input)
	if struct_error != nil {
		log.Println("Error 26: ", struct_error)

		fields := strings.Split(struct_error.Error(), ";")
		for idx, element := range fields {
			values := strings.Split(element, ":")
			if len(values) == 2 {
				tag := strings.TrimSpace(strings.ToLower(values[0]))
				msg := strings.TrimSpace(values[1])
				errors = append(errors, e.New(tag, idx+3, msg))
			}
		}
		return errors
	}
	return nil
}
