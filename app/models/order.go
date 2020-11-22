package models

import (
	"github.com/lib/pq"
	"time"
)

type Order struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	Student      Student `gorm:"foreignkey:StudentRefer"`
	StudentRefer uint64

	PaymentMethod string

	SubscriptionItem string

	Price float64

	Status string

	CreatedAt time.Time
}

func (o Order) Date() string {
	return o.CreatedAt.Format("02/01/2006")
}

func (o Order) Created() string {
	return o.CreatedAt.Format("02/01/2006 15:04:05")
}

type SubscriptionItem struct {
	Name         string
	Descriptions pq.StringArray `gorm:"type:varchar(200)[]"`
	Type         string
	Price        float64
	DayDuration  int64
}

func GetDonSubscription() SubscriptionItem {
	return SubscriptionItem{
		Name:         "Faire un Don",
		Descriptions: []string{"Accès à toutes les leçons", "Accès à tous les exercices", "Gagner des étoiles"},
		Type:         "don",
		Price:        4,
		DayDuration:  30,
	}
}


func GetSubscriptionItem(t string) (s SubscriptionItem) {
	switch t {
	case "Don":
		s = GetDonSubscription()
		break
	}

	return s
}
