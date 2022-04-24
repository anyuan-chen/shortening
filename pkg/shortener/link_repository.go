package shortener

type LinkRepository interface {
	CreateUser(id string) (User, error)
	Get(original_link string) (Link, error)
	Create(id string, shortened_link string, original_link string, user_id string) (Link, error)
	GetByUserID(user_id string)([]Link, error)
}