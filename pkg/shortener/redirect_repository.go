package shortener

type RedirectRepository interface {
	Get(shortened_link string) (string, error)  //returns the longer link
	Create(shortened_link string, original_link string, user_id string) error
	Delete(shortened_link string) error
}