package models

import (
	"time"
)

type Step struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	Title    string
	Subtitle string
	//TIPS: need to be unic in base
	Order         float64
	IsFree        bool
	IsChallenge   bool
	ChallengeName string

	IsDisplayHands bool

	//TIPS: Can be 'lesson_formula', 'lesson', 'write_number', 'check_hand', 'check_hand_formula', 'calcul', 'write_formula'
	Type string

	Lessons []Lesson `gorm:"foreignkey:StepRefer"`
	Trains  []Train  `gorm:"foreignkey:StepRefer"`

	Note int `gorm:"-"`

	CreatedAt time.Time
}

func (s Step) Created() string {
	return s.CreatedAt.Format("02/01/2006 15:04:05")
}

func (s Step) IsLesson() bool {
	return s.Type == "lesson"
}

func (s Step) IsWriteNumber() bool {
	return s.Type == "write_number"
}

func (s Step) IsCheckHand() bool {
	return s.Type == "check_hand"
}

func (s Step) IsCalcul() bool {
	return s.Type == "calcul"
}

func (s Step) GameType() string {
	gameType := ""

	switch s.Type {
	case "lesson_formula":
		gameType = "Formule par coeur"
		break
	case "lesson":
		gameType = "Leçon"
		break
	case "write_number":
		gameType = "Écrire les nombres"
		break
	case "check_hand":
		gameType = "Trouver les bonnes mains"
		break
	case "check_hand_formula":
		gameType = "Trouver la bonne formule"
		break
	case "calcul":
		gameType = "Calcul"
		break
	case "write_formula":
		gameType = "Remplir la formule"
		break
	}

	return gameType
}
