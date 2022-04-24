package shortener

type LinkService interface {
	Get(shortened_link string) (string, error)
	CreateAuthenticated(id string, shortened_linkstring , original_link string, user_id string) (Link, error)
	CreateUnauthenticated(id string, shortened_link string, original_link string)(Link, error)
	GetByUserID(session_id string)([]Link, error)
}

