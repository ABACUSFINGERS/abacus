package front

import (
	"../../models"
	"../../tools/debug"
	"../../tools/generate"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func AccountController(c echo.Context) error {
	debug.Log("front<AccountController>")

	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	student := models.StudentSession(db, c)
	if student.ID == 0 {
		return c.Redirect(http.StatusFound, "/connexion/")
	}
	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<account>", map[string]interface{}{
		"title":   "Mon compte",
		"student": student,
	})
}

func NewAccountController(c echo.Context) error {
	debug.Log("front<NewAccountController>")

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<_new_account>", map[string]interface{}{
		"title": "Nouveau compte",
	})
}

func PNewAccountController(c echo.Context) error {
	debug.Log("front<PAccountController>")
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	firsName := c.FormValue("firsName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	birthday, _ := time.Parse("2006-01-02", c.FormValue("birthday"))
	password := c.FormValue("password")

	student := models.Student{
		FirstName:        firsName,
		LastName:         lastName,
		Email:            email,
		Birthday:         birthday,
		Password:         password,
		EndLicence:       time.Now(),
		Token:            generate.Uuid(),
		Notes:            nil,
		LastConnectionAt: time.Now(),
		CreatedAt:        time.Now(),
	}

	db.Save(&student)

	if models.StudentLogin(db, c, email, password) {
		return c.Redirect(http.StatusFound, "/")
	} else {
		return c.Render(http.StatusOK, site.Host+"abacusfingers.local<home>", map[string]interface{}{
			"title": "Accueil",
		})
	}
}
