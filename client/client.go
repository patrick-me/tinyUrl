package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/patrick-me/tinyUrl/handlers"
	"io"
	"net/http"
	"time"
)

type TinyUrlClient interface {
	CreateShortURL(original string, expirationInHours time.Duration) (string, error)
}

type TinyUrlClientImpl struct {
	Client        *http.Client
	TinyURLOrigin string
}

func (c *TinyUrlClientImpl) CreateShortURL(original string, expirationInHours time.Duration) (string, error) {
	request := &handlers.URLRequest{Url: original, ExpirationInHours: expirationInHours}
	body, err := json.Marshal(request)

	if err != nil {
		fmt.Printf("Can't marshal url request: %s, %s\n", request, err)
		return "", err
	}

	req, err := http.NewRequest(
		"POST", c.TinyURLOrigin+"/short", bytes.NewBuffer(body),
	)

	resp, err := c.Client.Do(req)

	if err != nil {
		fmt.Printf("Error during get a response: %s\n", err)
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error during read a response body: %s\n", err)
		return "", err
	}

	if resp.StatusCode >= 400 {
		fmt.Printf("Error - status code: %d, %s, response body: %s\n", resp.StatusCode, resp.Status, respBody)
		return "", fmt.Errorf("error - status code: %d, %s", resp.StatusCode, resp.Status)
	}

	var urlResp handlers.URLResponse
	err = json.Unmarshal(respBody, &urlResp)

	if err != nil {
		fmt.Printf("Error during unmarshal a response body: %s\n", err)
		return "", err
	}

	return c.TinyURLOrigin + "/" + urlResp.Short, nil
}

func CreateTinyUrlClient(origin string) *TinyUrlClientImpl {
	return &TinyUrlClientImpl{
		TinyURLOrigin: origin,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		}}
}
