package response

import "time"

/**
* Response Type
 */

type Response struct {
	ForcePageRefresh bool        `json:"forcePageRefresh"`
	Message          string      `json:"message"`
	Status           bool        `json:"status"`
	Data             interface{} `json:"data"`
	Error            ErrorType   `json:"error"`
	StatusCode       int         `json:"statusCode"`
	Time             TimeType    `json:"time"`
}

type ErrorType struct {
	Exists bool   `json:"exists"`
	Errors string `json:"errors"`
}

type TimeType struct {
	UnixTime int32 `json:"unixTime"`
}

/**
* PARMS
* message
* status
* data
* error
* statuscode
 */
func ResponseWriter(message string, status bool, data interface{}, statusCode int) *Response {
	return &Response{
		false,
		message,
		status,
		data,
		ErrorType{false, ""},
		statusCode,
		TimeType{int32(time.Now().Unix())}}
}

func LoadErrorResponse(statusCode int, err error) *Response {
	var data interface{}
	data = "Oops!"
	return &Response{
		false,
		"Something went wrong. Please check error message for more information",
		false,
		data,
		ErrorType{true, err.Error()},
		statusCode,
		TimeType{int32(time.Now().Unix())}}
}
