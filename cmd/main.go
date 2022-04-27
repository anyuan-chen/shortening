package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anyuan-chen/urlshortener/server/pkg/api"
	"github.com/anyuan-chen/urlshortener/server/pkg/link_repository/cockroachdb"
	"github.com/anyuan-chen/urlshortener/server/pkg/redirect_repository/redis"
	"github.com/anyuan-chen/urlshortener/server/pkg/session_repository/inmemory"
	service "github.com/anyuan-chen/urlshortener/server/pkg/shortener/core_logic"
	"github.com/gorilla/mux"
)
type Server struct {
	r *mux.Router
}
func main() {
	r := mux.NewRouter()
    link_handler, err := cockroachdb.CreateCockroachDB(os.Getenv("COCKROACH_DB_DATABASE_URL"))
    if err != nil {
        log.Fatal("error creating link_handler" + err.Error())
    }
    redirect_handler, err := redis.CreateRedisRepository(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))
    if err != nil {
        log.Fatal("error creating redirect_handler"  + err.Error())
    }
    session_handler := &inmemory.MemorySessionRepository{}
    session_handler.CreateSessionRepository()
    service := service.NewLinkService(&redirect_handler, session_handler, &link_handler)
    api := api.NewService(service)
    http.HandleFunc("/login", api.Login)
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