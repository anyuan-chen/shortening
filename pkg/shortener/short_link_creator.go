package shortener

type ShortLinkCreator interface {
	GenerateShortLink(original_link string, user_id string) string //returns the short link
}