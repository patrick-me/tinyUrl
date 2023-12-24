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

#### Build


> docker build -t tiny-001 .


#### Run

> docker run -dp 127.0.0.1:8080:8080 tiny-001

