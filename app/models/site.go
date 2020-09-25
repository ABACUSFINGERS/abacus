package models

import (
	"time"
)

type Site struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	Host string
	Title  string
	Name     string

	CreatedAt        time.Time
}

func (s Site) Created() string {
	return s.CreatedAt.Format("02/01/2006 15:04:05")
}
