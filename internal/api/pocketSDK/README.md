
   * This is a SDK for the api of the pocket, which allows you to log in using tokens and add links to your pocket account.<br>
Example usage:
```go
func main() {
	//you can generate a consumer key on https://getpocket.com/developer/apps/
    u, err := pocket.NewClient("your consumerKey", "https://redirectUri.com")
    if err != nil {
        log.Fatal("new client")
    }

    req, err := u.GetRequestToken()
    if err != nil {
        log.Fatal(err)
    }

	// link must be sent to the client
    autorizatioLink, err := u.GetAutorizationUrl(req)
    if err != nil {
	log.Fatal(err)
    }

	//infoUser = access token && username
    infoUser, err := u.Authetication(req)
    if err != nil {
        log.Fatal(err)
    }

    u.AccessToken = infoUser.Access_token

	//now you can add links to your pocket
    err = u.Add("https://example.com")
    if err != nil {
        log.Fatal("add")
    }
}
```
