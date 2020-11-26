package dashboard

import (
	"mindset.go/abacus/app/models"
	"mindset.go/abacus/app/tools/debug"
	"github.com/labstack/echo"
	"net/http"
)

func HomeController(c echo.Context) error {
	debug.Log("dashboard<HomeController>")

	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	user := models.UserSession(db, c)
	if user.ID == 0 || user.IsAdmin == false {
		return c.Redirect(http.StatusFound, "/dashboard/login/")
	}

	return c.Render(http.StatusOK, "dashboard<home>", map[string]interface{}{
		"u":     user,
		"title": "Tableau de bord",
	})
}
