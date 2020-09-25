package models

import (
	"github.com/lib/pq"
	"time"
)

type Lesson struct {
	ID        uint64 `sql:"type:bigint PRIMARY KEY`
	StepRefer uint64
	Children  pq.StringArray `gorm:"type:varchar(200)[]"`

	CreatedAt time.Time
}

func (l Lesson) Created() string {
	return l.CreatedAt.Format("02/01/2006 15:04:05")
}
