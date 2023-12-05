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
	//вы можете сгенерировать consumer key на сайте https://getpocket.com/developer/apps/
    u, err := pocket.NewClient("your consumerKey", "https://redirectUri.com")
    if err != nil {
        log.Fatal("new client")
    }

	//запрашиваем request token
    req, err := u.GetRequestToken()
    if err != nil {
        log.Fatal(err)
    }

	// эту ссылку нужно отправить пользователю, чтобы он подтвердил право доступа вашего сервиса к его аккаунту
    autorizatioLink, err := u.GetAutorizationUrl(req)
    if err != nil {
	log.Fatal(err)
    }

	//обмением request token на access токен
    infoUser, err := u.Authetication(req)
    if err != nil {
        log.Fatal(err)
    }

    u.AccessToken = infoUser.Access_token

	//с помощью access токена добавляем ссылку в ваш pocket аккаунт
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
