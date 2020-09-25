package dashboard

import (
	"../../models"
	"../../tools/debug"
	"github.com/labstack/echo"
	"net/http"
)

func LoginController(c echo.Context) error {
	debug.Log("dashboard<LoginController>")

	return c.Render(http.StatusOK, "dashboard<_login>", map[string]interface{}{
		"title": "Identification",
	})
}

func PLoginController(c echo.Context) error {
	debug.Log("dashboard<PLoginController>")
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	email := c.FormValue("email")
	password := c.FormValue("password")

	if models.UserLogin(db, c, email, password) {
		return c.Redirect(http.StatusFound, "/dashboard/")
	} else {
		return c.Render(http.StatusOK, "dashboard<_login>", map[string]interface{}{
			"title": "Identification",
			"error": "Mauvais identifiant ou mot de passe.",
		})
	}
}

func DisconnectionController(c echo.Context) error {
	debug.Log("dashboard<DisconnectionController>")
	models.UserClearSession(c)
	return c.Redirect(http.StatusFound, "/dashboard/login/")
}
