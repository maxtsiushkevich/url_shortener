package models

import "time"

type URL struct {
	Code         string
	URL          string
	CreationTime time.Time
	Clicks       int
}

type ShortURLResponse struct {
	URL string
}
