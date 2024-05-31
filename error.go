package utils



type ErrorData struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
}

type ErrorService interface {
	GetErrorMessage(key string,statusCode int,lang string) error
	GetHttpStatusCode(err error) (string, int)
}