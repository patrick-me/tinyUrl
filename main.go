package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"tinyUrl/generators"
	"tinyUrl/handlers"
	"tinyUrl/storage"
	"tinyUrl/utils"
)

func main() {
	startApp()
}

func startApp() {
	appHandler := &handlers.UrlHandler{
		Storage: &storage.LocalMemoryStorage{
			Store: make(map[string]string),
		},
		Generator: &generators.SimpleRandGenerator{},
	}

	err := fasthttp.ListenAndServe(
		utils.GetEnv("APP_PORT", ":8080"),
		fasthttp.CompressHandler(appHandler.Router),
	)

	if err != nil {
		fmt.Println("App finished with err: ", err)
		return
	}
}
