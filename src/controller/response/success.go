package response

import (
	"net/http"
)

// Success - Send Success Result with the Object
func Success(body interface{}, w http.ResponseWriter) {
	writeResponse(200, body, w)
}
