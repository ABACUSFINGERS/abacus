package front

import (
	"mindset.go/abacus/app/models"
	"mindset.go/abacus/app/tools"
	"mindset.go/abacus/app/tools/debug"
	"github.com/labstack/echo"
	"net/http"
)

func CmsController(c echo.Context) error {
	debug.Log("front<CmsController>")

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

	var cms models.Cms
	db.Set("gorm:auto_preload", true).Where("slug = ?", tools.GetLastPathSegment(c)).First(&cms)

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<cms>", map[string]interface{}{
		"title":   site.Title,
		"content": cms.Content,
		"student": student,
	})
}
