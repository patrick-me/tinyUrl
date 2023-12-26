# tinyUrl
A simple tinyUrl

---

## Simple usage

#### Create URL

> curl -d '{"url":"https://ya.ru"}' localhost/short && echo
```
URL created!
"{\"url\":\"https://ya.ru\",\"short\":\"kc6fc\"}"
```

#### Get URL

> curl -L localhost/kc6fc
```
... Redirect to https://ya.ru
```

#### Default expiration
- 48 hours
---

## Docker usage

#### Compose
> docker-compose up
