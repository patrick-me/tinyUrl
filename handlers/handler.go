package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"tinyUrl/generators"
	"tinyUrl/storage"
)

type UrlHandler struct {
	Storage storage.Storage
}

func (h *UrlHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {

	if ctx.IsPost() {
		switch string(ctx.Path()) {
		case "/short":
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
		_, err = fmt.Fprintf(ctx,
			"URL not found!\n%q\n", url)

		if err != nil {
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

	if !isUrlValid(urlRequest.Url) {
		ctx.Error(fmt.Sprintf("URL should start with http(s)://, %q", urlRequest.Url), fasthttp.StatusBadRequest)
		return
	}

	shortURL := generators.GenRandURL()

	for h.Storage.Contains(shortURL) {
		shortURL = generators.GenRandURL()
	}
	h.Storage.Save(shortURL, urlRequest.Url)

	data, _ := json.Marshal(&URLResponse{Url: urlRequest.Url, Short: shortURL})

	ctx.SetStatusCode(fasthttp.StatusCreated)
	_, err = fmt.Fprintf(ctx,
		"URL created!\n%q\n", string(data))

	if err != nil {
		fmt.Println("Error code #123")
		return
	}
}

func isUrlValid(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}
