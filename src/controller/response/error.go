package response

import (
	"net/http"
)

// BadRequest - Send Bad Request with Message
func BadRequest(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(400, body, w)
}

// InternalServerError - Send Internal Server Error with Message
func InternalServerError(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(500, body, w)
}
