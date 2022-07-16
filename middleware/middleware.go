package middleware

import (
	"net/http"
	"github.com/MelvinRB27/server-user/authorization"
)

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if _, err := authorization.ValidateToken(token);
		err != nil {
			forbidden(w, r)
			return
		}
		
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Does not have authorization"))
}