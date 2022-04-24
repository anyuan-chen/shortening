package api

type Shortener interface {
	
}
//Redirect takes a URL which was the result of a previous shortening operation,
//then redirects the user to the original URL.
func Redirect(){

}
//CreateAuthenticated is meant as a way for logged in users to shorten a link.
//This will then be accessible to them if they use the GetLinksForUserID endpoint
func CreateAuthenticated(){

}
//CreateUnauthenticated is a meant as a way for unauthenticated users to shorten a link.
func CreateUnauthenticated(){

}
//GetLinksForUserID returns all links created by a specific user from the CreateAuthenticated
//handler.
func GetLinksForUserID(){

}