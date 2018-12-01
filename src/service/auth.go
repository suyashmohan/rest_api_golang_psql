package service

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

const superSecret = "secretABCD1234"

// AuthJWTClaims - Our Custom Claims for Auth
type AuthJWTClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// Valid - Check for Validation for JWT Token
func (c AuthJWTClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}

	if c.UserID == "" {
		return errors.New("Must provide a user ID")
	}

	return nil
}

// Auth - Check for Authorization
func Auth(h httprouter.Handle, errorHandle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) < 2 {
			errorHandle(w, r, ps)
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])

		token, _ := jwt.ParseWithClaims(reqToken, &AuthJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(superSecret), nil
		})

		if claims, ok := token.Claims.(*AuthJWTClaims); ok && token.Valid {
			// Extend Params by adding user id
			i := len(ps)
			psNew := make([]httprouter.Param, i+1)
			copy(psNew, ps)
			psNew[i].Key = "userid"
			psNew[i].Value = claims.UserID

			h(w, r, psNew)

		} else {
			errorHandle(w, r, ps)
		}
	}
}

// JWT - Generate a JWT Token
func JWT(userid int64) (string, error) {
	expiry := time.Now().Add(24 * time.Hour) // Expire after 1 day
	claims := AuthJWTClaims{
		strconv.FormatInt(userid, 10),
		jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(superSecret))
}
