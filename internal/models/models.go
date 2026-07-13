package models

import "time"

type Url struct {
	Code         string
	Url          string
	CreationTime time.Time
	Clicks       int
}
