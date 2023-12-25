# tinyUrl
A simple tinyUrl

---

## Simple usage

#### Create URL


> curl -d '{"url":"https://ya.ru"}' localhost:8080/short && echo
```
URL created!
"{\"url\":\"https://ya.ru\",\"short\":\"kc6fc\"}"
```

#### Get URL

> curl -L localhost:8080/kc6fc
```
... Redirect to https://ya.ru
```

---

## Docker usage

#### Compose
> docker-compose up
