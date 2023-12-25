package main

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
	"tinyUrl/db"
	"tinyUrl/generators"
	"tinyUrl/handlers"
	"tinyUrl/storage"
	"tinyUrl/utils"
)

func main() {
	startApp()
}

func startApp() {
	fmt.Println("Starting app")
	appHandler := &handlers.UrlHandler{
		//TODO: add config to choose storage
		/*Storage: &storage.LocalMemoryStorage{
			Store: make(map[string]string),
		},*/
		Storage: &storage.RedisStorage{
			Client:     db.CreateRedisClient(),
			Context:    context.Background(),
			Expiration: 48 * time.Hour,
		},
		Generator: &generators.SimpleRandGenerator{},
	}

	fmt.Println("Starting fasthttp")
	err := fasthttp.ListenAndServe(
		utils.GetEnv("APP_PORT", ":8080"),
		fasthttp.CompressHandler(appHandler.Router),
	)

	if err != nil {
		fmt.Println("App finished with err: ", err)
		return
	}
}
