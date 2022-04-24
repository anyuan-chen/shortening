package shortener

type linkService struct {
	redirectRepository RedirectRepository
	sessionRepository SessionRepository
	linkRepository LinkRepository
}