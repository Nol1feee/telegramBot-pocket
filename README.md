<div align="center">
    <h1>TG-BOT</h1>
    <h5>
        Телеграм бот, который сохраняет ваши ссылки в ваш <a href="https://getpocket.com/en/">pocket.</a>
    </h5>
    <p>
        Русский | <a href="README.en.md">English</a> 
    </p>
</div>

---

## О проекте и технологиях
Была использована сторонняя библиотека..
С помощью самописного SDK для <ссылка> api pocket'a реализовал логику тг бота, который помогает сохранять ссылки в аккаунт pocket'a.
Для аутентификации пользователя с помощью chi был поднят отдельный сервер, который ловил редирект от покета и запрашивал accesses token, который далее сохранялся в boltDB. Развертывался проект в docker'e.

## Установка
```
git clone https://github.com/Nol1feee/telegramBot-pocket
```
## Пример использования SDK
- с помощью SDK вы можете получить request token, обменять его на access token и добавлять ссылки в pocket аккаунт
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
## запуск проекта
1. настройте переменные окружение
```
export CONSURMER_KEY= <вы можете сгенерировать consumer key на сайте https://getpocket.com/developer/apps/>
export TGBOT_TOKEN= <зарегистрировать бота и сгененировать токен по следующей ссылке https://t.me/BotFather>
```
2. запустите докер контейнер
```
docker run -p <8888:8888> Dockerfile
```
