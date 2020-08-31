package res

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS = "SUCCESS"
	ERROR = "ERROR"

	UnsupportedEntity = "Unsupported entity."
	UnprocessableEntity = "Unprocessable entity."
)

// a basic/generic response struct
type response struct {
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	Status string `json:"status"`
	//ApplicationStatusCode int `json:"application_status_code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// any error response strcut
type errResponse struct {
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	//ApplicationStatusCode int `json:"application_status_code,omitempty"`
	Status string `json:"status"`
	Errors interface{} `json:"errors,omitempty"`
}

// represents any successful request
func Send(w http.ResponseWriter,message string, data interface{},code int) {
	sendResponse(w,response{
		Code:    code,
		Status: SUCCESS,
		Message: message,
		Data: data,
	},code)
}

// represents any error resonse
func SendError(w http.ResponseWriter, message string, errors interface{},code int) {
	sendResponse(w,errResponse{
		Code:    code,
		Status: ERROR,
		Message: message,
		Errors:  errors,
	},code)
}

// the actual function that sends the response
func sendResponse(w http.ResponseWriter,res interface{},code int) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}