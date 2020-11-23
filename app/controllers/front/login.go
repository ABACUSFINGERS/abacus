package front

import (
	"flycode.go/abacusf/app/models"
	"flycode.go/abacusf/app/tools/debug"
	"github.com/labstack/echo"
	"net/http"
)

func LoginController(c echo.Context) error {
	debug.Log("front<LoginController>")

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<_login>", map[string]interface{}{
		"title": "Identification",
	})
}

func PLoginController(c echo.Context) error {
	debug.Log("front<PLoginController>")
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	var site models.Site
	for _, s := range _sites {
		if s.Host == c.Request().Host {
			site = s
		}
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	if models.StudentLogin(db, c, email, password) {
		return c.Redirect(http.StatusFound, "/")
	} else {
		return c.Render(http.StatusOK, site.Host+"abacusfingers.local<_login>", map[string]interface{}{
			"title": "Identification",
			"error": "Mauvais identifiant ou mot de passe.",
		})
	}
}

func DisconnectionController(c echo.Context) error {
	debug.Log("front<DisconnectionController>")
	models.StudentClearSession(c)
	return c.Redirect(http.StatusFound, "/")
}
