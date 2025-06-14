package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/patrick-me/tinyUrl/generators"
	"github.com/patrick-me/tinyUrl/storage"
	"github.com/valyala/fasthttp"
	"strings"
)

const (
	APIV1ShortURL = "/api/v1/shortUrl"
)

type UrlHandler struct {
	Storage   storage.Storage
	Generator generators.URLGenerator
	TinyHost  string
}

func (h *UrlHandler) Router(ctx *fasthttp.RequestCtx) {

	if ctx.IsPost() {

		switch string(ctx.Path()) {
		case APIV1ShortURL:
			createShortUrlHandler(ctx, h)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
		return
	}

	if ctx.IsGet() {
		getHandler(ctx, h)
		return
	}

	ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
}

func getHandler(ctx *fasthttp.RequestCtx, h *UrlHandler) {
	url := string(ctx.Path()[1:]) // without leading slash
	originUrl, err := h.Storage.Get(url)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		_, err2 := fmt.Fprintf(ctx,
			"URL not found!\n%q\n", url)

		if err2 != nil {
			fmt.Println("Error code #122")
			return
		}
		return
	}

	ctx.Redirect(originUrl, fasthttp.StatusMovedPermanently)

}

func createShortUrlHandler(ctx *fasthttp.RequestCtx, h *UrlHandler) {
	body := ctx.PostBody()
	var urlRequest *URLRequest
	err := json.Unmarshal(body, &urlRequest)

	if err != nil {
		ctx.Error(fmt.Sprintf("Can't parse url, %q", err), fasthttp.StatusBadRequest)
		return
	}

	if valid, message := isUrlValid(urlRequest.Url); !valid {
		ctx.Error(message, fasthttp.StatusBadRequest)
		return
	}

	shortURL := h.Generator.GetRandURL()

	for h.Storage.Contains(shortURL) {
		shortURL = h.Generator.GetRandURL()
	}
	h.Storage.Save(shortURL, urlRequest.Url, urlRequest.ExpirationInHours)

	data, _ := json.Marshal(&URLResponse{Url: h.TinyHost + "/" + shortURL})

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody(data)
}

func isUrlValid(url string) (bool, string) {
	url = strings.ToLower(url)

	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		return false, fmt.Sprintf("URL should start with http(s)://, %q", url)
	}

	if strings.Contains(url, "localhost") || strings.Contains(url, "127.0.0.1") {
		return false, fmt.Sprintf("URL isn't valid, %q", url)
	}

	return true, ""
}
