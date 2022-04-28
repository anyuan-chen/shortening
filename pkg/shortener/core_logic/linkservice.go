package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/anyuan-chen/urlshortener/server/pkg/shortener"
	"golang.org/x/oauth2"
)

type LinkService struct {
	redirectRepository shortener.RedirectRepository
	sessionRepository shortener.SessionRepository
	linkRepository shortener.LinkRepository
	shortLinkCreator shortener.ShortLinkCreator
}

func NewLinkService (redirRepo shortener.RedirectRepository, sessRepo shortener.SessionRepository, linkRepo shortener.LinkRepository, shortLinkCreator shortener.ShortLinkCreator) *LinkService {
	return &LinkService{
		redirectRepository: redirRepo,
		sessionRepository: sessRepo,
		linkRepository: linkRepo,
		shortLinkCreator: shortLinkCreator,
	}
}

func (ls *LinkService) Get(shortened_link string) (string, error){
	cacheLink, cacheErr := ls.redirectRepository.Get(shortened_link)
	databaseLink, databaseErr := ls.linkRepository.Get(shortened_link)
	if cacheErr != nil && databaseErr != nil {
		return "", errors.New("link not found")
	}
	if cacheLink != ""{
		return cacheLink, nil
	}
	return databaseLink, nil
}

func (ls *LinkService) CreateAuthenticated(original_link string, user_id string) (shortener.Link, error){
	shortened_link := ls.shortLinkCreator.GenerateShortLink(original_link, user_id)
	link, err := ls.linkRepository.Create(shortened_link, original_link, user_id) 
	if err != nil {
		return shortener.Link{}, err
	}
	err = ls.redirectRepository.Create(shortened_link, original_link, user_id)
	if err != nil {
		err = errors.New("cache error - successfully saved into main database")
	}
	return link, err
}

func (ls *LinkService) CreateUnauthenticated(original_link string) (shortener.Link, error){
	shortened_link := ls.shortLinkCreator.GenerateShortLink(original_link, "guest")
	link, err := ls.linkRepository.Create(shortened_link, original_link, "guest") 
	if err != nil {
		return shortener.Link{}, err
	}
	err = ls.redirectRepository.Create(shortened_link, original_link, "guest")
	if err != nil {
		return shortener.Link{}, err
	}
	return link, err
}

func (ls *LinkService) GetByUserID(session_id string) ([]shortener.Link, error) {
	id, err := ls.sessionRepository.GetId(session_id);
	if err != nil {
		return nil, err
	}
	links, err := ls.linkRepository.GetByUserID(id);
	if err != nil {
		return nil, err
	}
	return links, nil 
}

func (ls *LinkService) Login (provider string, oauthstate string) (string, error) {
	url, err := ls.sessionRepository.GetLoginRedirect(provider, oauthstate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (ls *LinkService) Callback(provider string, code string) (*oauth2.Token, error) {
	fmt.Print("made it to callback")
	token, err := ls.sessionRepository.CodeExchange(provider, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (ls *LinkService) CreateSession(access_token string, refresh_token string, token_type string, expiry time.Time, provider string)(string, error){
	session_id, err := ls.sessionRepository.CreateSession(access_token, refresh_token, token_type, expiry, provider)
	if err != nil {
		return "", err
	}
	return session_id, nil
}
func (ls *LinkService) ValidateSession(session_id string) (string, error) {
	info, err := ls.sessionRepository.GetId(session_id)
	if err != nil {
		return "", err 
	}
	return info, nil
}
func (ls *LinkService) GetSession( session_id string)(shortener.Session, error){
	session, err := ls.sessionRepository.GetSession(session_id)
	if err != nil {
		return shortener.Session{}, err
	}
	return session, nil
}