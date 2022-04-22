package main

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/handler"
	"github.com/anyuan-chen/urlshortener/server/store"
	"github.com/anyuan-chen/urlshortener/server/users"
	"github.com/gorilla/mux"
)
type Server struct {
	r *mux.Router
}
func main() {
	
	r := mux.NewRouter()
	store.InitializeStore()
	users.InitializeDatabase()

	r.HandleFunc("/create/{url}", handler.CreateShortUrl).Methods("POST")
	r.HandleFunc("/redirect/{url}", handler.RedirectURL).Methods("GET")
	r.HandleFunc("/auth/google/login", handler.OauthGoogleLogin)
	r.HandleFunc("/auth/google/callback", handler.OauthGoogleCallback)
	r.HandleFunc("/auth/github/login", handler.OauthGithubLogin)
	r.HandleFunc("/auth/github/callback", handler.OauthGithubCallback)
	r.Handle("/id", handler.LoggedInMiddleware(handler.GetUser))
	r.Handle("/links", handler.LoggedInMiddleware(handler.GetLinksForUser))
	r.HandleFunc("/loggedin", handler.IsLoggedIn)
	http.Handle("/", &Server{r})
	// handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }
    // Lets Gorilla work
    s.r.ServeHTTP(rw, req)
}