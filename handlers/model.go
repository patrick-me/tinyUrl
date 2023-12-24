package handlers

type URLRequest struct {
	Url string `json:"url"`
}

type URLResponse struct {
	Url   string `json:"url"`
	Short string `json:"short"`
}
