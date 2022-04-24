package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
type Server struct {
	r *mux.Router
}
func main() {
	r := mux.NewRouter()
	http.Handle("/", &Server{r})
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
    if req.Method == "OPTIONS" {
        return
    }
    s.r.ServeHTTP(rw, req)
}