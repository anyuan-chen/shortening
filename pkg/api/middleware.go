package api

import (
	"context"
	"net/http"
)

//Authenticate is HTTP middleware that takes the next function, then attempts to verify if the user
//has a valid active session. If they are authenticated, the user_id and their session_id are passed on
//through context to the next HTTP handler function
func (s *Service ) Authenticate(next func (w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		session_id, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "bad session id", http.StatusBadRequest)
			return
		}	
		ctx := r.Context()
		user_id, err := s.linkService.ValidateSession(session_id.Value)
		if err != nil {
			http.Error(w, "no active session", http.StatusUnauthorized)
			return
		}
		type key string
		ctx = context.WithValue(ctx, key("id") , user_id)
		ctx = context.WithValue(ctx, key("session_id") , session_id)
		next(w, r.WithContext(ctx))
	})
}
