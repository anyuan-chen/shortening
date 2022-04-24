package shortener

type Link struct {
	Id string `json:"id"`
	Shortened_link string `json:"shortened_link"`
	Original_link string `json:"original_link"`
	User_id string `json:"user_id"`
}

type User struct {
	Id string `json:"id"`
}

type Session struct {
	Access_token string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
	Token_type string `json:"token_type"`
	Expiry string `json:"expiry"`
	Provider string `json:"provider"`
}