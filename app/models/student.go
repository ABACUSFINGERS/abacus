package models

import (
	"flycode.go/abacusf/app/tools/generate"
	session "github.com/ipfans/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

type Student struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	FirstName string
	LastName  string
	Email     string

	Birthday time.Time

	Password string
	Token    string

	EndLicence time.Time

	Notes []Note `gorm:"foreignkey:StudentRefer"`

	LastConnectionAt time.Time
	CreatedAt        time.Time
}

func (s Student) Created() string {
	return s.CreatedAt.Format("02/01/2006 15:04:05")
}

func (s Student) LastConnection() string {
	return s.LastConnectionAt.Format("02/01/2006 15:04:05")
}

func (s Student) EnLicenceDate() string {
	return s.EndLicence.Format("02/01/2006")
}

func (s Student) IsLicenced() bool {
	return s.EndLicence.After(time.Now())
}

func (s Student) Steps(db *gorm.DB) (steps []Step) {
	db.Order("order").Find(&steps)

	for i, step := range steps {
		for _, note := range s.Notes {
			if note.StepRefer == step.ID {
				steps[i].Note = note.Note
			}
		}
	}
	return steps
}

func StudentLogin(db *gorm.DB, c echo.Context, email string, password string) bool {
	var student Student

	db.Where("email LIKE ? AND password LIKE ?", email, password).First(&student)

	if student.ID != 0 {
		student.Token = generate.Uuid()
		db.Save(&student)

		sess := session.Default(c)
		sess.Set("s_uuid", student.Token)
		_ = sess.Save()

		return true
	} else {
		return false
	}
}

func StudentSession(db *gorm.DB, c echo.Context) Student {
	var student Student

	sess := session.Default(c)
	uuid := sess.Get("s_uuid")

	db.Where("token LIKE ?", uuid).First(&student)

	if student.ID != 0 {
		student.LastConnectionAt = time.Now()
		db.Save(&student)
	}

	return student
}

func StudentClearSession(c echo.Context) {
	sess := session.Default(c)
	sess.Set("s_uuid", "")
	_ = sess.Save()
}

func (s Student) IsConnected() bool {
	if s.ID != 0 {
		return true
	}
	return false
}

func (s Student) AddLicenceDays(db *gorm.DB, d int) {
	s.EndLicence = time.Now().AddDate(0, 0, d)

	db.Save(&s)
}