package middlewares

import (
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			err := utils.Authorizer(jwtToken)
			if err != nil {
				http.Error(w, utils.ErrTokenSignature.Error(), http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}
}
