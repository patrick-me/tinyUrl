package main

import (
	"github.com/valyala/fasthttp"
	"tinyUrl/handlers"
	"tinyUrl/storage"
)

func main() {
	myHandler := &handlers.UrlHandler{Storage: &storage.LocalMemoryStorage{Store: make(map[string]string)}}
	fasthttp.ListenAndServe(":8080", fasthttp.CompressHandler(myHandler.HandleFastHTTP))

}
