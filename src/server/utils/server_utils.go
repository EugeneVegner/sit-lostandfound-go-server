package utils

import (
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
	e "src/server/errors"
	//"io/ioutil"
	//"io"
	//"encoding/json"
	"github.com/asaskevich/govalidator"
	log "src/server/logger"
	"strings"
)

func EncodeBody(ctx *gin.Context, input interface{}) []e.Error {
	var errors []e.Error
	if err := ctx.BindJSON(&input); err != nil {
		errors = append(errors, e.New("body", e.ServerErrorInvalidBody, err.Error()))
		return errors
	}

	// Validate input structure
	return ValidateEntity(ctx, input)
}

func ValidateEntity(ctx *gin.Context, entity interface{}) []e.Error {
	log.Func(ValidateEntity)
	var errors []e.Error
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		log.Debug("invalidate: ", err)
		fields := strings.Split(err.Error(), ";")
		for _, element := range fields {
			values := strings.Split(element, ":")
			if len(values) == 2 {
				tag := strings.TrimSpace(strings.ToLower(values[0]))
				msg := strings.TrimSpace(values[1])
				errors = append(errors, e.New(tag, e.ServerErrorInvalidStructure, msg))
			}
		}
		return errors
	}
	return nil
}

