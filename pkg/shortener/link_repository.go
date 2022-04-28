package shortener

type LinkRepository interface {
	CreateUser(id string) (error)
	Get(shortened_link string) (string, error)
	Create(shortened_link string, original_link string, user_id string) (Link, error) //for id in table, maybe implement some sort of hash
	GetByUserID(user_id string)([]Link, error)
	DeleteUser(user_id string)(error)
	DeleteLink(id string)(error)
	GetUser(user_id string) (string, error)
}