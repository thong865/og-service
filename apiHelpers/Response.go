package apihelpers

import (
	"encoding/json"
	"net/http"
)

//ResponseData structure
type ResponseData struct {
	Error   bool        `json:"error"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

//Message returns map data
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond returns basic response structure
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
