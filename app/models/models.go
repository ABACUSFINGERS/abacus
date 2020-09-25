package models

import (
	"../tools/debug"
	"github.com/jinzhu/gorm"
	"os"
)

func GetDataBase() (db *gorm.DB) {
	coStr := "host=" + os.Getenv("DATABASE_HOST") +
		" port=" + os.Getenv("DATABASE_PORT") +
		" user=" + os.Getenv("DATABASE_USER") +
		" password=" + os.Getenv("DATABASE_PASSWORD") +
		" dbname=" + os.Getenv("DATABASE_NAME") +
		" sslmode=disable"
	db, err := gorm.Open("postgres", coStr)
	if err != nil {
		panic(err)
	}

	debug.Log(">%v : db instance create\n", db)
	return db
}

func Migrate() {
	db := GetDataBase()
	db.AutoMigrate(
		&Site{},
		&User{},
		&Cms{},
		&Student{},
		&Lesson{},
		&Train{},
		&TrainChild{},
		&Step{},
		&Note{},
		&Order{},
	)

	debug.Log("<%v : db instance close\n", db)
	db.Close()
}

