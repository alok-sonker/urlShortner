package main

import "time"

type URLJSON struct {
	URL string `json:"url"`
}

type Server struct {
}
type URLMetadata struct {
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
	Counter   int32     `json:"counter"`
}
