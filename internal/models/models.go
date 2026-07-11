package models

import "time"

type URL struct {
	Url          string
	CreationTime time.Time
	Clicks       int
}
