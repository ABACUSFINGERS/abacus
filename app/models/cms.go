package models

import "html/template"

type Cms struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	Content template.HTML
	Slug    string `gorm:"unique"`
}
