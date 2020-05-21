package functions

import (
	"bytes"
	"encoding/json"
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

//ResponseLoginJSON function
//Function that formats response to JSON in google callback
func ResponseLoginJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

//PrettyJSONTerminal function
//Function that beautifully prints a json object
func PrettyJSONTerminal(data []byte) {
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, data, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(dst.String())
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
