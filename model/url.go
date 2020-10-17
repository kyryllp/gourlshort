package model

type URL struct {
	ID           int    `json:"id"`
	RedirectName string `json:"redirect_name"`
	OriginalUrl  string `json:"original_url"`
}
