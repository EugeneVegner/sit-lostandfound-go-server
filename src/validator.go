package src

//import (
//	"net/http"
//	"encoding/json"
//	"github.com/asaskevich/govalidator"
//	"strings"
//	"io/ioutil"
//	"io"
//	"log"
//)

//type ValidateResult struct {
//	Errors	[]Error
//	Success bool
//}
//
//func (result *ValidateResult) AddError(error Error) []Error  {
//	result.Errors = append(result.Errors, error)
//	return result.Errors
//}
//
//func ValidateBody(w http.ResponseWriter, r *http.Request, input interface{}) ValidateResult {
//
//	var result ValidateResult
//
//	body, err := ioutil.ReadAll(r.Body)
//	switch {
//	case err == io.EOF:
//		log.Println("Empty body: ")
//		result.Errors = append(result.Errors, MakeError("body",21, "Request body is empty"))
//		result.Success = false
//		return result
//
//	case err != nil:
//		log.Println("Error body: ", err)
//		result.Errors = append(result.Errors, MakeError("body",22, err.Error()))
//		result.Success = false
//		return result
//	}
//
//	//err := json.NewDecoder(r.Body).Decode(input)
//	if parser_error := json.Unmarshal(body, input); parser_error != nil {
//		log.Println("Error 24: ", parser_error)
//		result.Errors = append(result.Errors, MakeError("body", 23, parser_error.Error()))
//		result.Success = false
//		return result
//	}
//
//	// Validate input structure
//	_, struct_error := govalidator.ValidateStruct(input)
//	if struct_error != nil {
//		log.Println("Error 26: ", struct_error)
//
//		fields := strings.Split(struct_error.Error(), ";")
//		for _, element := range fields {
//			values := strings.Split(element, ":")
//			if len(values) == 2 {
//				tag := strings.TrimSpace(strings.ToLower(values[0]))
//				msg := strings.TrimSpace(values[1])
//				result.AddError(MakeError(tag, 26, msg))
//			}
//		}
//
//		result.Success = false
//		return result
//
//	}
//
//	result.Success = true
//	return result
//}





