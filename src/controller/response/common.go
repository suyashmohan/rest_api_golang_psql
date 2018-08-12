package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeResponse(status int, body interface{}, w http.ResponseWriter) {
	json, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", json)
}

func createBodyWithMessage(message string) interface{} {
	body := map[string]string{
		"message": message,
	}
	return body
}
