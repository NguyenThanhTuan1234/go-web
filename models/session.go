package models

import (
	"time"
)

type Session struct {
	Id int
	// Un           string
	LastActivity time.Time
}
