package handlers

import "time"

type URLRequest struct {
	Url               string        `json:"url"`
	ExpirationInHours time.Duration `json:"expirationInHours"`
}

type URLResponse struct {
	Url   string `json:"url"`
	Short string `json:"short"`
}
