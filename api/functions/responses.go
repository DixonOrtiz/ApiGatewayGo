package functions

import (
	"fmt"
	"net/http"
)

//ResponseJSON function
//Function that formats all server responses to JSON
func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}

//ERROR function
//Function that response an error from the server
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ResponseJSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	ResponseJSON(w, http.StatusBadRequest, nil)
}
