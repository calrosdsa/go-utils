package error

import (
    _r "github.com/calrosdsa/go-utils"
	"errors"
	"net/http"

	"github.com/goccy/go-json"
)

type errorService struct {
	locale _r.Locale
}

func NewService(locale _r.Locale) _r.ErrorService{
	return &errorService{
		locale: locale,
	}
}

func(s *errorService) GetErrorMessage(key string,statusCode int,lang string) (err error){
	message := s.locale.MustLocalize(key,lang)
	data := _r.ErrorData{
		Message: message,
		StatusCode: statusCode,
	}
	j,_ := json.Marshal(data)
	return errors.New(string(j))
}

func (s *errorService) GetHttpStatusCode(err error)(message string,statusCode int){
	var data _r.ErrorData
	err1 := json.Unmarshal([]byte(err.Error()),&data)
	if err1!= nil {
		return err.Error(),http.StatusBadRequest
	}
	return data.Message,data.StatusCode
}
