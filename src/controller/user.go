package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"../repository"
	"./request"
	"./response"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// UserController - Controller for User API
type UserController struct {
	UserRepo *repository.UserRepository
}

// CreateUser - Create a new User in DB
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userReq := request.NewUserRequest{}
	json.NewDecoder(r.Body).Decode(&userReq)

	if len(strings.TrimSpace(userReq.Username)) == 0 || len(strings.TrimSpace(userReq.Password)) == 0 {
		response.BadRequest("Incorrect format for Username/Password", w)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.MinCost)
	if err != nil {
		response.InternalServerError("Something wrong went with Password", w)
		return
	}
	user := uc.UserRepo.New(userReq.Username, string(hash))

	if user != nil {
		response.Success(user, w)
	} else {
		response.BadRequest("Unable to create a new User with the given username/password", w)
	}
}
