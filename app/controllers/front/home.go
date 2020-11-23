package front

import (
	"flycode.go/abacusf/app/models"
	"flycode.go/abacusf/app/tools/debug"
	"github.com/labstack/echo"
	"net/http"
)

func HomeController(c echo.Context) error {
	debug.Log("front<HomeController>")

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

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<home>", map[string]interface{}{
		"title":   site.Title,
		"student": student,
	})
}
