package models

import (
	"mindset.go/abacus/app/tools/generate"
	session "github.com/ipfans/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

type User struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	FirstName string
	LastName  string
	Email     string
	Phone     string

	Password string
	Token    string

	IsAdmin bool

	LastConnectionAt time.Time
	CreatedAt        time.Time
}

func (u User) Created() string {
	return u.CreatedAt.Format("02/01/2006 15:04:05")
}

func (u User) LastConnection() string {
	return u.LastConnectionAt.Format("02/01/2006 15:04:05")
}

func UserLogin(db *gorm.DB, c echo.Context, email string, password string) bool {
	var user User

	db.Where("email LIKE ? AND password LIKE ?", email, password).First(&user)

	if user.ID != 0 {
		user.Token = generate.Uuid()
		db.Save(&user)

		sess := session.Default(c)
		sess.Set("d_uuid", user.Token)
		_ = sess.Save()

		return true
	} else {
		return false
	}
}

func UserSession(db *gorm.DB, c echo.Context) User {
	var user User

	sess := session.Default(c)
	uuid := sess.Get("d_uuid")

	db.Where("token LIKE ?", uuid).First(&user)

	if user.ID != 0 {
		user.LastConnectionAt = time.Now()
		db.Save(&user)
	}

	return user
}

func UserClearSession(c echo.Context) {
	sess := session.Default(c)
	sess.Set("d_uuid", "")
	_ = sess.Save()
}
