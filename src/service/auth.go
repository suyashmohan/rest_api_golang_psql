package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt"

	"github.com/julienschmidt/httprouter"
)

const superSecret = "secretABCD1234"

func getSigner() jwt.Signer {
	return jwt.HS256(superSecret)
}

// Auth - Check for Authorization
func Auth(h httprouter.Handle, errorHandle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		jwtReq, err := jwt.FromRequest(r)
		if err != nil {
			errorHandle(w, r, ps)
			return
		}
		if err := jwtReq.Verify(getSigner()); err != nil {
			errorHandle(w, r, ps)
			return
		}
		userID := jwtReq.Public()["userid"]
		if userID, ok := userID.(float64); ok {
			// Extend Params by adding user id
			i := len(ps)
			psNew := make([]httprouter.Param, i+1)
			copy(psNew, ps)
			psNew[i].Key = "userid"
			psNew[i].Value = strconv.FormatInt(int64(userID), 10) // JWT Returns number as float64 but it's integer and Param needs it as string :(

			h(w, r, psNew)
		} else {
			h(w, r, ps)
		}
	}
}

// JWT - Generate a JWT Token
func JWT(userid int64) (string, error) {
	expiry := time.Now().Add(24 * time.Hour) // Expire after 1 day
	payload := make(map[string]interface{})
	payload["userid"] = userid

	return jwt.Sign(getSigner(), &jwt.Options{ExpirationTime: expiry, Public: payload})
}
