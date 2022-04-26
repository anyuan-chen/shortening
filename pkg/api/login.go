//Package api provides a HTTP API with authentication and link shortening capabilities.
package api

import "github.com/anyuan-chen/urlshortener/server/pkg/shortener"

type service struct {
	linkService shortener.LinkService
}

func NewService(linkService shortener.LinkService) service{
	return service{linkService: linkService}
}
func Login(){
	
}
func Callback(){

}