FROM golang:1.21-alpine

WORKDIR /tinyurl

COPY generators ./generators
COPY handlers ./handlers
COPY storage ./storage
COPY main.go ./
COPY go.mod ./

RUN go mod download
RUN go mod tidy

RUN go build -o /tiny

CMD [ "/tiny" ]

EXPOSE 8080
