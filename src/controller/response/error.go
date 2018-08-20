package response

import (
	"net/http"
)

// BadRequest - Send Bad Request with Message
func BadRequest(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(400, body, w)
}

// Unauthorized - Send Unauthorized Error with Message
func Unauthorized(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(401, body, w)
}

// Forbidden - Send Forbidden Error with Message
func Forbidden(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(403, body, w)
}

// NotFound - Send NotFound Error with Message
func NotFound(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(404, body, w)
}

// InternalServerError - Send Internal Server Error with Message
func InternalServerError(message string, w http.ResponseWriter) {
	body := createBodyWithMessage(message)
	writeResponse(500, body, w)
}
