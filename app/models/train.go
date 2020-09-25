package models

import (
	"regexp"
	"time"
)

type Train struct {
	ID        uint64 `sql:"type:bigint PRIMARY KEY`
	StepRefer uint64

	//TIPS: C'est la liste des opérations (1+2, 1, 50-1+3)
	Children []TrainChild `gorm:"foreignkey:TrainRefer"`

	CreatedAt time.Time
}

type TrainChild struct {
	ID         uint64 `sql:"type:bigint PRIMARY KEY`
	TrainRefer uint64

	Operation string
	Result    string

	//TIPS: Temps en millisecondes
	Speed int64

	CreatedAt time.Time
}

func (t Train) Created() string {
	return t.CreatedAt.Format("02/01/2006 15:04:05")
}

//TODO: Vérifier que cette fonction fonctionne bien
// Est censé retourner : ['3', '+', '7']
func (t TrainChild) ChildrenSplit() []string {
	re := regexp.MustCompile(`[+-]`)
	childSplit := re.Split(t.Operation, -1)

	return childSplit
}
