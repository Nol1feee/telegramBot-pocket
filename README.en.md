<div align="center">
    <h1>TG-BOT</h1>
    <h5>
        A Telegram bot that saves your links to your <a href="https://getpocket.com/en/">pocket.</a>
    </h5>
    <p>
        Русский | <a href="README.en.md">English</a> 
    </p>
</div>

---

## About the project and technology
A third-party library was used.
With the help of self-written SDK for <link> api pocket'a implemented the logic of tg bot, which helps to save links to pocket'a account.
For user authentication with chi was raised a separate server, which caught redirect from the pocket and requested accesses token, which is further stored in boltDB. The project was deployed in docker.

## Installation
```
git clone https://github.com/Nol1feee/telegramBot-pocket
```
## Example of use SDK
- with SDK you can get request token, exchange it for access token and add links to pocket account
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
## project launch
1. configure .env variables
```
export CONSURMER_KEY= <вы можете сгенерировать consumer key на сайте https://getpocket.com/developer/apps/>
export TGBOT_TOKEN= <зарегистрировать бота и сгененировать токен по следующей ссылке https://t.me/BotFather>
```
2. run the docker container
```
docker run -p <8888:8888> Dockerfile
```
