package main

import (
	"../models"
	"../tools"
	"../tools/debug"
	"../tools/generate"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	_ = godotenv.Load(".env")

	models.Migrate()
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	InitializeSites(db)
	InitializeStudents(db)
	InitializeGames(db)
	InitializeCmss(db)
	InitializeSounds()
}

func InitializeSites(db *gorm.DB) {
	var count int64
	db.Table("sites").Count(&count)
	if count == 0 {
		site := models.Site{
			ID:        1,
			Host:      "abacusfingers.website",
			Title:     "AbacusFingers la nouvelle méthode d'apprentissage mathématique !",
			Name:      "AbacusFingers",
			CreatedAt: time.Time{},
		}
		db.Save(&site)
	}
}

func InitializeStudents(db *gorm.DB) {
	var count int64
	db.Table("students").Count(&count)
	if count == 0 {
		user1 := models.Student{
			FirstName:        "Salah",
			LastName:         "Eldin",
			Email:            "contact@abacusfingers.com",
			Birthday:         time.Time{},
			Password:         "0000",
			Token:            generate.Uuid(),
			EndLicence:       time.Now().AddDate(100, 0, 0),
			Notes:            nil,
			LastConnectionAt: time.Time{},
			CreatedAt:        time.Time{},
		}

		db.Save(&user1)

		user2 := models.Student{
			FirstName:        "Chahine",
			LastName:         "Eldin",
			Email:            "abacusfingers@gmail.com",
			Birthday:         time.Time{},
			Password:         "0000",
			Token:            generate.Uuid(),
			EndLicence:       time.Now().AddDate(100, 0, 0),
			Notes:            nil,
			LastConnectionAt: time.Time{},
			CreatedAt:        time.Time{},
		}

		db.Save(&user2)
	}
}

func InitializeGames(db *gorm.DB) {
	var count int64
	db.Table("steps").Count(&count)
	if count == 0 {

		//STEP 1
		lesson := models.Lesson{
			Children:  []string{"0", "1", "2", "3", "4", "5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step := models.Step{
			Title:          "Tutoriel des nombres de 0 à 5",
			Subtitle:       "Faire défiler l'image des doigts et des chiffres de 0 à 5.",
			Order:          1,
			IsFree:         true,
			IsDisplayHands: true,
			Type:           "lesson",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 2
		lesson = models.Lesson{
			Children:  []string{"0", "1", "2", "3", "4", "5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Reconnaître les nombres de 0 à 5",
			Subtitle:       "Affichage des doigts de 0 à 5 puis saisir la bonne réponse.",
			Order:          2,
			IsFree:         true,
			IsDisplayHands: true,
			Type:           "write_number",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 3
		lesson = models.Lesson{
			Children:  []string{"4", "3", "5", "2", "0", "1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Trouver les nombres de 0 à 5",
			Subtitle:       "Affichage des chiffres de 0 à 5 puis sélectionner la bonne image.",
			Order:          3,
			IsFree:         true,
			IsDisplayHands: true,
			Type:           "check_hand",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 4
		step = models.Step{
			Title:          "Addition de 0 à 5",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          4,
			IsFree:         true,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train := models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild := models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+2",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+1",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+2",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+3",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3+1",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3+2",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4+1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 5
		lesson = models.Lesson{
			Children:  []string{"6", "7", "8", "9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Tutoriel des nombres de 6 à 9",
			Subtitle:       "Faire défiler l'image des doigts et des chiffres de 6 à 9. Une voix énonce les chiffres afin de les reconnaitres.",
			Order:          5,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "lesson",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 6
		lesson = models.Lesson{
			Children:  []string{"6", "7", "8", "9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Reconnaître les nombres de 6 à 9",
			Subtitle:       "Affichage des doigts de 6 à 9 puis saisir la bonne réponse.",
			Order:          6,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "write_number",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 7
		lesson = models.Lesson{
			Children:  []string{"6", "7", "8", "9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Trouver les nombres de 6 à 9",
			Subtitle:       "Affichage des chiffres de 6 à 9 puis sélectionner la bonne image.",
			Order:          7,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "check_hand",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 8
		step = models.Step{
			Title:          "Addition de 5 à 9",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          8,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+1",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+2",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+3",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+4",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+1",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+2",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+3",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+1",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+2",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+1",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 9
		step = models.Step{
			Title:          "Soustraction de 0 à 4",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          9,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-1",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-2",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-3",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-4",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-2",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-3",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1-1",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 10
		step = models.Step{
			Title:          "Soustraction de 5 à 9",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          10,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9-4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9-3",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9-2",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9-1",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8-3",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8-2",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8-1",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7-2",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7-1",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6-1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//?STEP Challenge
		// step = models.Step{
		// 	Title:          "Soustraction et Addition",
		// 	Subtitle:       "Défilement des chiffres 1 par 1 mode flash.",
		// 	Order:          10.1,
		// 	IsFree:         false,
		// 	IsChallenge:    true,
		// 	ChallengeName:  "Master YOCTO",
		// 	IsDisplayHands: true,
		// 	Type:           "calcul",
		// 	Note:           0,
		// 	CreatedAt:      time.Time{},
		// }

		// db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+1-1+2-1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2-1-1+3-2",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3+1-2-1+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-1+1-2+1",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3-4+3-3",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9-2-2+1+3",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8-3+2+2-4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7-2+4-3+1",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+1+3-4+3",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+2-3+4-4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 11
		lesson = models.Lesson{
			Children:  []string{"10", "20", "30", "40", "50"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Tutoriel des nombres de 0 à 5",
			Subtitle:       "Faire défiler l'image des doigts et des chiffres de 10 à 50. Une voix énonce les chiffres afin de les reconnaitres.",
			Order:          11,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "lesson",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 12
		lesson = models.Lesson{
			Children:  []string{"10", "20", "30", "40", "50"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Reconnaître les nombres de 10 à 50",
			Subtitle:       "Affichage des doigts de 10 à 50 puis saisir la bonne réponse.",
			Order:          12,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "write_number",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 13
		lesson = models.Lesson{
			Children:  []string{"10", "20", "30", "40", "50"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Trouver les nombres de 10 à 50",
			Subtitle:       "Affichage des chiffres de 10 à 50 puis sélectionner la bonne image.",
			Order:          13,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "check_hand",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 14
		step = models.Step{
			Title:          "Addition de 10 à 50",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          14,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+10",
			Result:     "20",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+20",
			Result:     "30",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+30",
			Result:     "40",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+40",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20+10",
			Result:     "30",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20+20",
			Result:     "40",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20+30",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30+10",
			Result:     "40",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30+20",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40+10",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 15
		lesson = models.Lesson{
			Children:  []string{"60", "70", "80", "90"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Tutoriel des nombres de 60 à 90",
			Subtitle:       "Faire défiler l'image des doigts et des chiffres de 60 à 90.",
			Order:          15,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "lesson",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 16
		lesson = models.Lesson{
			Children:  []string{"60", "70", "80", "90"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Reconnaître les nombres de 60 à 90",
			Subtitle:       "Affichage des doigts de 60 à 90 puis saisir la bonne réponse.",
			Order:          16,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "write_number",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 17
		lesson = models.Lesson{
			Children:  []string{"60", "70", "80", "90"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Trouver les nombres de 60 à 90",
			Subtitle:       "Affichage des chiffres de 60 à 90 puis sélectionner la bonne image.",
			Order:          17,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "check_hand",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 18
		step = models.Step{
			Title:          "Addition de 60 à 90",
			Subtitle:       "Défilement des chiffres 1 par 1 mode flash afin de s'initier à l'apprentissage de la méthode.",
			Order:          18,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+10",
			Result:     "60",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+20",
			Result:     "70",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+30",
			Result:     "80",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+40",
			Result:     "90",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60+10",
			Result:     "70",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60+20",
			Result:     "80",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60+30",
			Result:     "90",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70+10",
			Result:     "80",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70+20",
			Result:     "90",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80+10",
			Result:     "90",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 19
		lesson = models.Lesson{
			Children:  []string{"58", "42", "19", "32", "89", "5", "99", "51", "70", "12"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Reconnaître les nombres de 1 à 99",
			Subtitle:       "Affichage des doigts de 1 à 99 puis saisir la bonne réponse.",
			Order:          19,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "write_number",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 20
		lesson = models.Lesson{
			Children:  []string{"46", "20", "17", "11", "67", "23", "97", "32", "79", "12"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		step = models.Step{
			Title:          "Trouver les nombres de 1 à 99",
			Subtitle:       "Affichage des chiffres de 1 à 99 puis sélectionner la bonne image.",
			Order:          20,
			IsFree:         false,
			IsDisplayHands: true,
			Type:           "check_hand",
			Lessons:        []models.Lesson{lesson},
			Trains:         nil,
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		//STEP 21
		step = models.Step{
			Title:          "Addition de 11 à 99",
			Subtitle:       "Défilement des chiffres mode flash 2 chiffres en défilement.",
			Order:          21,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+1",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20+2",
			Result:     "22",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30+9",
			Result:     "39",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40+8",
			Result:     "48",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+7",
			Result:     "57",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60+5",
			Result:     "65",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70+2",
			Result:     "72",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80+3",
			Result:     "83",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90+1",
			Result:     "91",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 22
		step = models.Step{
			Title:          "Soustraction avec 2 chiffres",
			Subtitle:       "Défilement des chiffres mode flash 2 chiffres en défilement.",
			Order:          22,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40-10",
			Result:     "30",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40-20",
			Result:     "20",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40-30",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "40-40",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30-10",
			Result:     "20",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30-20",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30-30",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20-10",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20-20",
			Result:     "0",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 23
		step = models.Step{
			Title:          "Soustraction avec 2 chiffres",
			Subtitle:       "Défilement des chiffres mode flash 2 chiffres en défilement.",
			Order:          23,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90-10",
			Result:     "80",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90-20",
			Result:     "70",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90-30",
			Result:     "60",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90-40",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80-10",
			Result:     "70",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80-20",
			Result:     "60",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80-30",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70-10",
			Result:     "60",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70-20",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60-10",
			Result:     "50",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 24
		step = models.Step{
			Title:          "1ère étape vers la RAPIDITÉ",
			Subtitle:       "Les 4 formules des Petits Amis (+)",
			Order:          24,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+5-4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+2", "+5-3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+3", "+5-2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+4", "+5-1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4(+1)",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3(+2)",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2(+3)",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1(+4)",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 25
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          25,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+4", "+5-1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+4",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3+4",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4+4",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 26
		step = models.Step{
			Title:          "Exercices sur les Petits Amis +",
			Subtitle:       "Petit Ami (+) la formule",
			Order:          26,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "write_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+5-4",
			Result:     "+1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+5-3",
			Result:     "+2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+5-2",
			Result:     "+3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+5-1",
			Result:     "+4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		//STEP 27
		step = models.Step{
			Title:          "Exercices sur les Petits Amis",
			Subtitle:       "Exercices sur les Petits Amis",
			Order:          27,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+4-1-3+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+2-2+1+4",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-2-1+3+4",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-3-1+4+4",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3-4+2+3",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-2-2+3+3",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3-4+3+4",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-2-1+3+2",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+1-3+4+2",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-2-2+4+1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 28
		step = models.Step{
			Title:          "2ième étape vers la RAPIDITÉ.",
			Subtitle:       "Les Petits Amis (-) se composent en 4 formules.",
			Order:          28,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-1", "-5+4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-2", "-5+3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-3", "-5+2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-4", "-5+1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(-1)",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(-2)",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(-3)",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(-4)",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 29
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          29,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-4", "-5+1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5-4",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6-4",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7-4",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8-4",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 30
			step = models.Step{
					Title:          "Exercices sur les Petits Amis +",
					Subtitle:       "Petit Ami (+) la formule",
					Order:          30,
					IsFree:         false,
					IsDisplayHands: false,
					Type:           "write_formula",
					Note:           0,
					CreatedAt:      time.Time{},
				}
		
				db.Save(&step)
		
				train = models.Train{
					StepRefer: step.ID,
					CreatedAt: time.Time{},
				}
		
				db.Save(&train)
		
				trainChild = models.TrainChild{
					TrainRefer: train.ID,
					Operation:  "-5+4",
					Result:     "-1",
					Speed:      0,
					CreatedAt:  time.Time{},
				}
		
				db.Save(&trainChild)
		
				trainChild = models.TrainChild{
					TrainRefer: train.ID,
					Operation:  "-5+3",
					Result:     "-2",
					Speed:      0,
					CreatedAt:  time.Time{},
				}
		
				db.Save(&trainChild)
		
				trainChild = models.TrainChild{
					TrainRefer: train.ID,
					Operation:  "-5+2",
					Result:     "-3",
					Speed:      0,
					CreatedAt:  time.Time{},
				}
		
				db.Save(&trainChild)
		
				trainChild = models.TrainChild{
					TrainRefer: train.ID,
					Operation:  "-5+1",
					Result:     "-4",
					Speed:      0,
					CreatedAt:  time.Time{},
				}
		
				db.Save(&trainChild)
		
				lesson = models.Lesson{
					StepRefer: step.ID,
					Children:  []string{"+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5"},
					CreatedAt: time.Time{},
				}
		
				db.Save(&lesson)
		
		//STEP 31
		step = models.Step{
			Title:          "Exercices sur les Petits Amis (-)",
			Subtitle:       "Exercices sur les Petits Amis (-)",
			Order:          31,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5-4+1+2+3",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5-3-2+1+4",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5-2-1+3+4",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5-1-4+4+4",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3-4+2+3",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-2-2+3+3",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "1+3-4+3+4",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "3-2-1+3+2",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "2+1-3+4+2",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "4-2-2+4+1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 32
		step = models.Step{
			Title:          "3ième étape vers la RAPIDITÉ",
			Subtitle:       "Les 9 formules des bons Amis (+)",
			Order:          32,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+10-9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+2", "+10-8"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+3", "+10-7"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+4", "+10-6"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+5", "+10-5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+6", "+10-4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+7", "+10-3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+8", "+10-2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+9", "+10-1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+1)",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+2)",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+3)",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+4)",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+5)",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+6)",
			Result:     "15",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+7)",
			Result:     "16",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+8)",
			Result:     "17",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9(+9)",
			Result:     "18",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 33
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          33,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "10-9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+10-9",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)


		//STEP 34
		step = models.Step{
			Title:          "Exercices sur les Bon Amis (+)",
			Subtitle:       "Bon Ami (+) la formule",
			Order:          34,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "write_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-9",
			Result:     "+1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-8",
			Result:     "+2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-7",
			Result:     "+3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-6",
			Result:     "+4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-5",
			Result:     "+5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-4",
			Result:     "+6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-3",
			Result:     "+7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-2",
			Result:     "+8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-1",
			Result:     "+9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		//STEP 35
		step = models.Step{
			Title:          "Exercices sur les bons Amis",
			Subtitle:       "Exercices sur les bons Amis",
			Order:          35,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+1+4-4+4",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+2+3-4+3",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+2+4-4+2",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+3+2-4+1",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+3+3-4+4",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+3+4-4+3",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+4+1-4+2",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+4+2-4+1",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+4+3-4+4",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+4+4-4+3",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+5-4+4-2",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+5+1-4+1",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+5+2-4+4",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+5+3-4+3",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+5+4-4+2",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+6+4-4+1",
			Result:     "16",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+7+3-4+4",
			Result:     "19",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+7+4-4+3",
			Result:     "18",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+8+2-4+2",
			Result:     "17",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+8+3-4+1",
			Result:     "16",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+8+4-4+4",
			Result:     "19",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "9+9+1-4+3",
			Result:     "18",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+9+1-4+3",
			Result:     "18",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+9+2-4+2",
			Result:     "17",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+9+3-4+1",
			Result:     "16",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+9+4-4+4",
			Result:     "19",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 36
		step = models.Step{
			Title:          "4ème étape vers la RAPIDITÉ",
			Subtitle:       "Les 9 formules des bons Amis (-)",
			Order:          36,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-1", "-10+9"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-2", "-10+8"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-3", "-10+7"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-4", "-10+6"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-5", "-10+5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-6", "-10+4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-7", "-10+3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-8", "-10+2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-9", "-10+1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-1)",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-2)",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-3)",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-4)",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-5)",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-6)",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-7)",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-8)",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10(-9)",
			Result:     "1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 37
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          37,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-6", "-10+4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-6",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 38
		step = models.Step{
			Title:          "Exercices sur les Bons Amis (-)",
			Subtitle:       "Bons Ami (-) la formule",
			Order:          38,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "write_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+9",
			Result:     "-1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+8",
			Result:     "-2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+7",
			Result:     "-3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+6",
			Result:     "-4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+5",
			Result:     "-5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+4",
			Result:     "-6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+7",
			Result:     "-3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+8",
			Result:     "-2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-10+9",
			Result:     "-1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		//STEP 39
		step = models.Step{
			Title:          "Exercices sur les bons Amis (-)",
			Subtitle:       "Exercices sur les bons Amis (-)",
			Order:          39,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-1-4-4+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-2-3-4+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-2-4-4+2",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-3-3-4+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-3-4-4+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-4-1-4+2",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-4-2-4+1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-4-3-4+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-4-4-4+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-5+4-5-2",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-5-1-4+1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-5-1-4+1",
			Result:     "2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-5-2-4+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-5-3-4+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-5-4-4+2",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-6-4+4+1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-7-3+4+4",
			Result:     "19",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-7-4+4+3",
			Result:     "7",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-8-2+4+2",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-8-2+4+1",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-8-4+6+4",
			Result:     "10",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10-9-1+5+3",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-9-2+6+2",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-9-3+7+1",
			Result:     "8",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-9-4+5+4",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 40
		step = models.Step{
			Title:          "5ème étape vers la RAPIDITÉ",
			Subtitle:       "Les 4 formules des grands Amis (+)",
			Order:          40,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+6", "+10-5+1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+7", "+10-5+2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+8", "+10-5+3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+9", "+10-5+4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+9)",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+8)",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+7)",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+6)",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 41
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          41,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+6", "5+10-5+1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+6",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+6",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+6",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+6",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		
		//STEP 42
		step = models.Step{
			Title:          "Exercices sur les Grands Amis (+)",
			Subtitle:       "Grands Ami (+) la formule",
			Order:          42,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "write_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "+10-5+1",
			Result:     "+6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+3",
			Result:     "-2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+2",
			Result:     "-3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+1",
			Result:     "-4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+2", "+3", "+4", "+5","+6","+7","+8","+9","+10", "-1", "-2", "-3", "-4", "-5", "-6", "-7", "-8", "-9", "-10"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		//STEP 43
		step = models.Step{
			Title:          "Exercices sur les grands Amis (-)",
			Subtitle:       "Exercices sur les grands Amis (-)",
			Order:          43,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+6-4+1+4",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+6-3-3+3",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+6-4-2+2",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "8+6-2-4+1",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+7-3-4+4",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+7-4-1+3",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "7+7-1-3+2",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+8-2+1+1",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "6+8-3+2+4",
			Result:     "17",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5+9-4+2+3",
			Result:     "15",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 44
		step = models.Step{
			Title:          "6ème étape vers la RAPIDITÉ",
			Subtitle:       "Les 4 formules des grands Amis (+)",
			Order:          44,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "lesson_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-6", "-10+5-1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-7", "-10+5-2"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-8", "-10+5-3"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-9", "-10+5-4"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+9)",
			Result:     "14",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+8)",
			Result:     "13",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)
	
		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+7)",
			Result:     "12",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "5(+6)",
			Result:     "11",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 45
		step = models.Step{
			Title:          "Sélectionner la réponse la plus rapide.",
			Subtitle:       "Afin que l'enfant prenne conscience de l'efficacité des formules.",
			Order:          45,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "check_hand_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"-6", "-10+5-1"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-10+5-1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

	
		//STEP 46
		step = models.Step{
			Title:          "Exercices sur les Petits Amis +",
			Subtitle:       "Petit Ami (+) la formule",
			Order:          46,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "write_formula",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+4",
			Result:     "-1",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+3",
			Result:     "-2",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+2",
			Result:     "-3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "-5+1",
			Result:     "-4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		lesson = models.Lesson{
			StepRefer: step.ID,
			Children:  []string{"+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5"},
			CreatedAt: time.Time{},
		}

		db.Save(&lesson)

		//STEP 47
		step = models.Step{
			Title:          "Exercices sur les grands Amis (-)",
			Subtitle:       "Exercices sur les grands Amis (-)",
			Order:          47,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "11-6-4+1+4",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-6-3-3+3",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-6-4-2+2",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-6-2-4+1",
			Result:     "3",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "12-7-3-4+4",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-7-4-1+3",
			Result:     "4",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-7-1-3+2",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "13-8-2+1+1",
			Result:     "5",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-8-3+2+4",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-8-3+2+4",
			Result:     "9",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "14-9-4+2+3",
			Result:     "6",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 48
		step = models.Step{
			Title:          "Introduction du calcul au-delà de 99.",
			Subtitle:       "Introduction du calcul au-delà de 99.",
			Order:          48,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+100+100+100+100",
			Result:     "500",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "200+200+200+200+200",
			Result:     "1000",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "300+100+200+100+100",
			Result:     "800",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "200+400+100+100+100",
			Result:     "900",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+500+100+200+100",
			Result:     "1000",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "600+100+100+100+100",
			Result:     "1000",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+200+200+200+300",
			Result:     "1000",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+300+100+200+200",
			Result:     "900",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+200+100+300+100",
			Result:     "800",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "400+200+100+100+100",
			Result:     "900",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 49
		step = models.Step{
			Title:          "Uniquement les dixaines mais somme au-delà de 99",
			Subtitle:       "Uniquement les dixaines mais somme au-delà de 99",
			Order:          49,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+50+50+50+50",
			Result:     "250",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+20+30+40+50",
			Result:     "150",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+40+30+20+10",
			Result:     "150",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "20+20+30+30+30",
			Result:     "130",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+40+30+20+10",
			Result:     "150",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "70+60+50+40+30",
			Result:     "250",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "30+40+50+60+70",
			Result:     "250",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "80+90+60+50+40",
			Result:     "300",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90+40+50+20+50",
			Result:     "250",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "90+80+50+50+70",
			Result:     "340",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		//STEP 50
		step = models.Step{
			Title:          "Uniquement les dixaines mais somme au-delà de 99",
			Subtitle:       "Uniquement les dixaines mais somme au-delà de 99",
			Order:          50,
			IsFree:         false,
			IsDisplayHands: false,
			Type:           "calcul",
			Note:           0,
			CreatedAt:      time.Time{},
		}

		db.Save(&step)

		train = models.Train{
			StepRefer: step.ID,
			CreatedAt: time.Time{},
		}

		db.Save(&train)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "50+100+150+60+100",
			Result:     "460",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "10+30+200+400+250",
			Result:     "890",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "300+100+250+230+120",
			Result:     "1000",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "220+400+110+100+150",
			Result:     "980",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "700+10+30+50+60",
			Result:     "850",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "800+50+50+30+60",
			Result:     "990",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "900+50+10+10+20",
			Result:     "250",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "60+50+60+70+350",
			Result:     "590",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "180+200+350+200+50",
			Result:     "980",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)

		trainChild = models.TrainChild{
			TrainRefer: train.ID,
			Operation:  "100+550+100+60+50",
			Result:     "860",
			Speed:      0,
			CreatedAt:  time.Time{},
		}

		db.Save(&trainChild)
	}
}

func InitializeCmss(db *gorm.DB) {
	var count int64
	db.Table("cms").Count(&count)
	if count == 0 {
		c := models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>A propos</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column\"> <h2>Mémoire, concentration et calcul mental</h2> <p> ABACUS Fingers propose une approche alternative pour l’apprentissage des mathématiques et plus particulièrement le calcul mental. La méthode traditionnelle qui est enseigné à nos enfants consiste à comprendre la technique de l’opération alors qu’avec d’autres méthodes reposant sur les abaques ou sur les mains on travaille une gymnastique gestuelle. On apprend ainsi à mémoriser des gestes. Cette méthode permet grâce à une pratique régulière de devenir un pro du calcul mental et de se réconcilier avec les maths de manière générale. Les supports numériques font partie de notre quotidien, même l’institution scolaire s’est mise en marche vers les supports d’éducation numérique : les tableaux, les livres, les QR codes… Alors pourquoi pas utiliser ce support à bon escient et proposer un jeu pour stimuler l’intelligence de nos enfants. </p> <p> L’application ABACUS Fingers est née de cette réflexion : « Mettre à profil le support numérique pour proposer de l’inventivité pédagogique au service de l’enfant.\" ABACUS Fingers est une application d’entrainement dédié au calcul mental. Elle a été développée avec une approche ludique pour permettre aux enfants dès l’âge de 3 ans de se familiariser avec les chiffres et de s’entrainer au dénombrement mathématique. Comme son nom l’indique, notre application ABACUS Fingers repose sur l’utilisation de « ses propres mains » comme support de calcul. Grâce à cette gymnastique cérébrale régulière, l’enfant développe une agilité mentale et très vite accède à des raisonnements plus complexes. </p> Savoir manipuler les quantités devient utile pour structurer son cerveau, entraîner sa mémoire et améliorer son esprit d'analyse et de synthèse. </p> <p> « Un enfant qui calcul vite et un enfant qui a confiance en lui. Un enfant qui a confiance en lui est un enfant qui réussit. » </p> <p> Notre méthode l’apprentissage repose sur des étapes clefs qui doivent être respectés dans l’ordre proposés. Il est important de ne pas brûler ces étapes et de respecter le rythme de chaque enfant, rien ne sert de courir… De bien s’assurer que chaque étape est bien maitrisée avant de passer à la suivante. En effet, tel un jeu de dominos, chaque étape d’apprentissage permet de construire l’étape suivante. </p> </div> <div class=\"dt-sc-hr\"></div> <h2 class=\"dt-sc-hr-green-title\"> Nos croyances :</h2> <div class=\"column\"> <div class=\"column\"> <p> Nous pensons qu’il est impératif d’accompagner nos enfants vers le chemin de la réussite et de leur démontrer qu’ils sont tous capables de grandes choses. Que le succès réside dans la méthode et l’entrainement et que le champ du possible est immense... </p> <p> Dans ce sens, peu importe son niveau de départ, il suffit de s'y mettre. « C'est une gymnastique qui doit se pratiquer régulièrement, comme pour un instrument de musique ». </p> <p>Les postulats d’ABACUS Fingers : </p> <ul class=\"dt-sc-fancy-list star\"> <li>Le génie n’est pas un héritage</li> <li>Tous les enfants sont intelligents</li> <li>L’intelligence est acquise et non innée</li> <li>Avoir la « bosse des maths » est une théorie totalement infondée</li> </ul> </div> </div> </section> </div> </div>",
			Slug:    "a-propos",
		}
		db.Create(&c)
		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>Comment ça marche</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column\"> <p> La proposition d’ABACUS Fingers est de vous faire acquérir sa méthode de l’apprentissage des mathématiques par le biais du jeu.Vous pouvez découvrir les 5 premiers niveaux accessibles à tous gratuitement et sans profil. </p> <p> Pour accéder à la version complète du jeu, la création d’un profil est nécessaire. Celui-ci vous permettra de continuer à progresser et de reprendre au niveau auquel vous vous êtes arrêtés quelque soit le support que vous utilisez (ordinateur, tablette ou mobile). </p> <p> Une adresse email (email d’un parent) est nécessaire pour la création du profil.Il vous sera également demandé d’indiquer le nom de l’enfant, sa date de naissance, un mot de passe et de choisir une formule. </p> <p> Retrouve toutes les formules dans notre menu <a href=\"/cms/nos-tarifs\">«Nos Tarifs»</a>. </p> </div> </section> </div> </div>",
			Slug:    "comment-ca-marche",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"><div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>Nos tarifs</h1> </div> </div><div class=\"container\"><section id=\"primary\" class=\"content-full-width\"> <h2>Une formule adaptée à chacun</h2> <div class=\"dt-sc-pricing-table\"> <div class=\"column dt-sc-one-third first\"> <div class=\"dt-sc-pr-tb-col mustard\"> <div class=\"dt-sc-tb-header\"> <div class=\"dt-sc-tb-thumb\"><img src=\"http://placehold.it/300x200\" alt=\"\" title=\"\"> <div class=\"dt-sc-tb-title\"> <h3>Mensuel</h3></div> </div> <div class=\"dt-sc-price\"> 4 €<span> TTC</span></div> </div> <ul class=\"dt-sc-tb-content\"> <li>Accès à toutes les leçons</li> <li>Accès à tous les exercices</li> <li>Gagner des étoiles</li> </ul> <div class=\"dt-sc-buy-now\"><a href=\"/commander/?order_type=monthly\" class=\"dt-sc-button small mustard\">Commander</a></div> </div> </div> <div class=\"column dt-sc-one-third\"> <div class=\"dt-sc-pr-tb-col green\"> <div class=\"dt-sc-tb-header\"> <div class=\"dt-sc-tb-thumb\"><img src=\"http://placehold.it/300x200\" alt=\"\" title=\"\"> <div class=\"dt-sc-tb-title\"> <h3>Annuel</h3></div> </div> <div class=\"dt-sc-price\"> 20 €<span> TTC</span></div> </div> <ul class=\"dt-sc-tb-content\"> <li>Accès à toutes les leçons</li> <li>Accès à tous les exercices</li> <li>Gagner des étoiles</li> </ul> <div class=\"dt-sc-buy-now\"><a href=\"/commander/?order_type=annual\" class=\"dt-sc-button small green\">Commander</a></div> </div> </div> <div class=\"column dt-sc-one-third\"> <div class=\"dt-sc-pr-tb-col blue\"> <div class=\"dt-sc-tb-header\"> <div class=\"dt-sc-tb-thumb\"><img src=\"http://placehold.it/300x200\" alt=\"\" title=\"\"> <div class=\"dt-sc-tb-title\"> <h3>Bi-Annuel</h3></div> </div> <div class=\"dt-sc-price\"> 30 €<span> TTC</span></div> </div> <ul class=\"dt-sc-tb-content\"> <li>Accès à toutes les leçons</li> <li>Accès à tous les exercices</li> <li>Gagner des étoiles</li> </ul> <div class=\"dt-sc-buy-now\"><a href=\"/commander/?order_type=bi-annual\" class=\"dt-sc-button small blue\">Commander</a></div> </div> </div> </div> </section> <!--primary ends--> </div> <!--container ends--> </div>",
			Slug:    "nos-tarifs",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\">         	<!--breadcrumb-section starts-->             <div class=\"breadcrumb-section\">             	<div class=\"container\">                 	<h1>Blog Two Column</h1>                     <div class=\"breadcrumb\">                         <a href=\"index.html\">Home</a>                         <span class=\"fa fa-angle-double-right\"></span>                         <span class=\"current\">Blog Two Column</span>                     </div>                 </div>             </div>             <!--breadcrumb-section ends-->             <!--container starts-->             <div class=\"container\">             	<!--primary starts-->             	<section id=\"primary\" class=\"content-full-width\">                                   	<div class=\"column dt-sc-one-half first\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Activities Improves Mind </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                          <div class=\"column dt-sc-one-half\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Weekly Reader Zone </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                          <div class=\"column dt-sc-one-half first\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Wonderland Adventures </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                          <div class=\"column dt-sc-one-half\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Kids Innovations </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                          <div class=\"column dt-sc-one-half first\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Newspaper Activity </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                          <div class=\"column dt-sc-one-half\">                     	<article class=\"blog-entry\">                             <div class=\"blog-entry-inner\">                                 <div class=\"entry-meta\">                                     <a href=\"blog-detail.html\" class=\"blog-author\"><img src=\"http://placehold.it/90x90\" alt=\"\" title=\"\"></a>                                     <div class=\"date\">                                         <span> 27 </span>                                          <p> Aug <br> 2014 </p>                                      </div>                                     <a href=\"#\" class=\"comments\">                                         12 <span class=\"fa fa-comment\"> </span>                                     </a>                                     <a href=\"#\" class=\"entry_format\"><span class=\"fa fa-picture-o\"></span></a>	                                 </div>		                                 <div class=\"entry-thumb\">                                     <a href=\"blog-detail.html\"><img src=\"http://placehold.it/1170x711\" alt=\"\" title=\"\"></a>                                 </div>		                                 <div class=\"entry-details\">	                                     <div class=\"entry-title\">                                         <h3><a href=\"blog-detail.html\"> Co-curricular Activites </a></h3>                                     </div>			                                     <!--entry-metadata ends-->	                                     <div class=\"entry-body\">                                         <p>Roin bibendum nibhsds. Nuncsdsd fermdada msit ametadasd consequat. Praes porr nulla sit amet dui lobortis, id venenatis nibh accums. Proin lobortis tempus odio eget venenatis. Proin fermentum ut massa at bibendum. Proin bibendum non est quis egestas. Pellentesque at enim id enim tempus placerat. Etiam tempus gravida leo, et gravida justo bibendum non. Suspendisse vitae fermentum sapien.</p>                                     </div>	 		                                     <a href=\"blog-detail.html\" class=\"dt-sc-button small\"> Read More <span class=\"fa fa-chevron-circle-right\"> </span></a>		                                 </div>	                             </div>                         </article>                 	</div>                                    </section>                 <!--primary ends-->             </div>             <!--container ends-->         </div>",
			Slug:    "blog",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>La Magie des doigts</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column mb-4\"> <p>Abacus Fingers est un projet innovant qui vient en complément des méthodes traditionnelles de l’apprentissage des mathématiques. à travers notre plateforme multilingue, nous rendons accessible une approche du dénombrement très populaire en Asie mais encore méconnue dans le monde occidental. A travers des jeux de calcul mental, l’enfant se familiarise avec les mathématiques en utilisant la « mémoire de ses mains » particulièrement puissante chez lui, pour créer et consolider la numération puis le calcul. <br/><br/> Les résultats obtenus sont éloquents : tous les enfants progressent très sensiblement et définitivement. Ils deviennent des génies du calcul mental très jeunes et avec une simplicité déconcertante. Afin de sans rendre compte voici une illustration de la méthode et de ces <a class=\"text-violet\" href=\"https://youtu.be/fxGEkx-yrQI\" target=\"_blank\">formidables résultats</a>. </p></div> </section> </div> </div>",
			Slug:    "la-magie-des-doigts",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>La GAMIFICATION - Apprenez en jouant</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column mb-4\"> <p> Notre approche à travers notre plateforme est d’utiliser les codes des jeux vidéos. C’est ce qu’on appelle l’apprentissage par la gamification. La gamification repose sur différents principes dont la gratification par les récompenses immédiates. Elle permet de capter et de garder l’attention de l’enfant. Cette attention qui compose l’ingrédient fondamental de l’apprentissage est exprimée par la formule de la « Connaissance » : <br/><br/> <b style=\"font-size: 16px\">K = A x t</b> <br/> <b style=\"font-size: 16px\">Connaissance = Attention x temps</b> </p> </div> </section> </div> </div>",
			Slug:    "gamification",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>Le secret de la méthode</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column mb-4\"> <p> <u>Les sens au service de l'apprentissage :</u> <br/> Les neurosciences ont permis de démontrer que dans l’apprentissage, l’utilisation de plusieurs canaux sensoriels (nos yeux, nos oreilles, notre langue, notre nez et notre peau) améliore la compréhension et la mémorisation. <br/><br/> La force de la méthode ABACUS Fingers c’est quelle sollicite plusieurs nos sens pour l’apprentissage : </p> <div class=\"column\"> <ul class=\"dt-sc-fancy-list pink star\"> <li>Le visuel (pour l’affichage des chiffres)</li> <li>L’auditif (avec l’écoute des chiffres)</li> <li>Et le kinesthésique (pour les doigts)</li> </ul> </div> </div> </section> </div> </div>",
			Slug:    "secret-de-la-methode",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>L’histoire des ABACUS</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column\"> <p>En introduction, intéressons-nous au mot «CALCUL» qui vient du latin calculus (pluriel calculi) qui veut dire «caillou». Dans les temps anciens, on utilisait donc les cailloux pour procédé au dénombrement comme le bétail, les denrées alimentaires et plus largement les biens de la vie courant. De ce postulat, on introduira des cailloux de taille ou de forme distincte en leur attribuant des valeurs différentes afin d’éviter de manipuler un trop grand nombre de cailloux. Enfin, on songea à différencier la valeur des cailloux, non par leur taille ou leur forme, mais par leur emplacement sur un support plat organisée à cet effet. Il enrésulta la naissance de l’ABACUS qui dans la symbolique correspond aux premières calculatrices manuelles de l’époque.Concernant le terme ABACUS, il vient du latin qui désigne tous les instruments mécaniques plats dans lesquels on place des petits objets identiques (cailloux, jetons, boules...) sur des colonnes, des lignes ou des tiges parallèles, de manière à représenter des nombres à calculer.</p> <p><b>Types d’ABACUS</b></p> <p>Abacus de l’époque Romaine :</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/roman_abacus.jpg\"></div> <p>Le <b>Suanpan</b> de Chine :</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/suanpan.jpg\"></div> <p>Le <b>Soroban</b> du Japon :</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/abcus_yellow.jpg\"></div> <p>Le <b>Stichioty</b> d'origine Russe :</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/stchioty.jpg\"></div> <p>Le <b>Boulier</b> occidental :</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/boulier_occidental.jpg\"></div> </div> </section> </div> </div>",
			Slug:    "histoire-des-abacus",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>L’origine de la méthode</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column\"><p>La méthode ABACUS Fingers résulte d’une évolution de la méthode Japonaise qui utilise le SOROBAN. Le Soroban appartient à la famille des ABACUS. Il est apparut à la fin du 19 siècle au Japon. Il est encore très pratiqué au Japon et plus largement dans l’Asie.</p> <div class=\"text-center\"><img src=\"/assets/abacus/images/abcus_yellow.jpg\"></div> <p> L’approche d’ABACUS Fingers est de se dispenser du Soroban au profit des doigts et de la mémoire. Retrouvez-nous dans le menu «Entrer dans le JEU» afin de découvrir la méthode pas à pas d’ABACUS Fingers. </p> </div> </section> </div> </div>",
			Slug:    "origine-de-la-methode",
		}
		db.Create(&c)

		c = models.Cms{
			Content: "<div id=\"main\"> <div class=\"breadcrumb-section\"> <div class=\"container\"> <h1>L’histoire des ABACUS</h1> </div> </div> <div class=\"container\"> <section id=\"primary\" class=\"content-full-width\"> <div class=\"column\"> <p>Il existe différents championnats à travers le monde. Le dernier <a href=\"http://www.juniormentalcalculators.com/jmcwc-2019-competition/\" target=\"_blank\">«JUNIOR MENTALCALCULATION WORLD CHAMPIONSHIP »</a>, c’est déroulé en Allemagne le 03 octobre au 05 octobre 2019. Le concours se déroule en trois groupes d'âge :</p> <ul class=\"dt-sc-fancy-list star\"> <li>Juniors 1: Les moins de 11 ans</li> <li>Juniors 2: de 12 à 14 ans</li> <li>Seniors: de 15 à 19 ans</li> </ul> <p>Dans notre blog, nous ferons mentions de tous les championnats futurs des différents pays.</p> <p>Voici la liste de quelques grandes figures du calcul mental :</p> <ul class=\"dt-sc-fancy-list graduation-cap\"> <li>Alberto Coto García</li><li>Alexander Craig Aitken</li><li>Alexis Lemaire</li><li>Arthur F. Griffith</li><li>Carl Friedrich Gauss</li><li>Charles Grandemange</li><li>Daniel McCartney</li><li>Daniel Tammet</li><li>George Parker Bidder</li><li>Gottfried Ruckle</li><li>Hans Eberstark</li><li>Henri Mondeux</li><li>Herbert De Grote</li><li>Jacques InaudiJedediah Buxton</li><li>Kim PeekMohamed Yassine Chermiti</li><li>Pericles Diamandi</li><li>Peter M Deshong</li><li>Rudiger GrammSalo Finkelstein</li><li>Scott Flansburg</li><li>Shakuntala DeviThomas Fuller</li><li>Truman Henry Safford</li><li>Willis Dysart Wim KleinZacharias Dase</li><li>Zerah Colburn</li> </ul> </div> </section> </div> </div>",
			Slug:    "championnats-du-calcul-mental",
		}
		db.Create(&c)
	}
}



func InitializeSounds() {
	sounds := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "20", "30", "40", "50", "60", "70", "80", "90"}
	for _, sound := range sounds {
		tools.GetMP3SoundFromText(sound, "fr")
	}
}


