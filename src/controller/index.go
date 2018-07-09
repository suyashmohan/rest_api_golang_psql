package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexController - Controller for Index Routes
type IndexController struct{}

// IndexRoute - Route for Index ( / )
func (ic *IndexController) IndexRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{ \"message\": \"Hello! World\" }")
}
