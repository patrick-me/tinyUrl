# tinyUrl
A simple tinyUrl

---

## Simple usage

#### Create URL

> curl -d '{"url":"https://ya.ru"}' localhost/api/v1/shortUrl && echo
```
URL created!
{"shortUrl":"http://localhost/Ch9Lgr9kzW"}

```

#### Get URL

> curl -L http://localhost/Ch9Lgr9kzW
```
... Redirect to https://ya.ru
```

#### Default expiration
- 720 hours
---

## Docker usage

#### Compose
> docker-compose up

To rebuild use this cmd
> docker-compose -f docker-compose.yml up -d --build
