FROM golang:1.25-alpine AS build

WORKDIR /src

# With a vendor/ directory in the context Go builds offline (-mod=vendor is
# picked up automatically); without it the modules are fetched per go.sum.
COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/tiny .


FROM alpine:3.22

RUN adduser -D -u 10001 tiny

COPY --from=build /out/tiny /usr/local/bin/tiny

USER tiny

EXPOSE 8080

CMD [ "tiny" ]
