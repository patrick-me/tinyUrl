package main

import (
	"context"
	"fmt"
	"github.com/patrick-me/tinyUrl/db"
	"github.com/patrick-me/tinyUrl/generators"
	"github.com/patrick-me/tinyUrl/handlers"
	"github.com/patrick-me/tinyUrl/storage"
	"github.com/patrick-me/tinyUrl/utils"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	startApp()
}

func startApp() {
	fmt.Println("Starting app")
	expiration := utils.GetEnvInt("DEFAULT_STORAGE_EXPIRATION_HOURS", 720)
	tinyhost := utils.GetEnv("TINY_HOST", "http://localhost")

	appHandler := &handlers.UrlHandler{
		//TODO: add config to choose storage
		/*Storage: &storage.LocalMemoryStorage{
			Store: make(map[string]string),
		},*/

		Storage: &storage.RedisStorage{
			Client:     db.CreateRedisClient(),
			Context:    context.Background(),
			Expiration: time.Duration(expiration) * time.Hour,
		},
		Generator: &generators.SimpleRandGenerator{},
		TinyHost:  tinyhost,
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
