//Package api provides a HTTP API with authentication and link shortening capabilities.
package api

import (
	"net/http"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
)

type Service struct {
	linkService shortener.LinkService
}

func NewService(linkService shortener.LinkService) Service{
	return Service{linkService: linkService}
}

//Login is meant as an HTTP endpoint for users to login into the platform.
//This endpoint redirects to the user-parameter specified OAuth endpoint.
func (s *Service ) Login(w http.ResponseWriter, r *http.Request){
	
}

//Callback is an HTTP endpoint for the OAuth providers once they have logged
//into the platform. 
func (s *Service ) Callback(w http.ResponseWriter, r *http.Request){

}