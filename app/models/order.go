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

func GetMonthlySubscription() SubscriptionItem {
	return SubscriptionItem{
		Name:         "Mensuel",
		Descriptions: []string{"Accès à toutes les leçons", "Accès à tous les exercices", "Gagner des étoiles"},
		Type:         "monthly",
		Price:        4,
		DayDuration:  30,
	}
}

func GetAnnualSubscription() SubscriptionItem {
	return SubscriptionItem{
		Name:         "Annuel",
		Descriptions: []string{"Accès à toutes les leçons", "Accès à tous les exercices", "Gagner des étoiles"},
		Type:         "annual",
		Price:        20,
		DayDuration:  365,
	}
}

func GetBiAnnualSubscription() SubscriptionItem {
	return SubscriptionItem{
		Name:         "Bi-annuel",
		Descriptions: []string{"Accès à toutes les leçons", "Accès à tous les exercices", "Gagner des étoiles"},
		Type:         "bi-annual",
		Price:        30,
		DayDuration:  730,
	}
}

func GetSubscriptionItem(t string) (s SubscriptionItem) {
	switch t {
	case "monthly":
		s = GetMonthlySubscription()
		break
	case "annual":
		s = GetAnnualSubscription()
		break
	case "bi-annual":
		s = GetBiAnnualSubscription()
		break
	}

	return s
}
