package models

import (
	"time"
)

type Note struct {
	ID           uint64 `sql:"type:bigint PRIMARY KEY`
	StudentRefer uint64
	StepRefer    uint64

	Note int

	CreatedAt time.Time
}

func (n Note) Created() string {
	return n.CreatedAt.Format("02/01/2006 15:04:05")
}
