package front

import (
	"mindset.go/abacus/app/models"
	"mindset.go/abacus/app/tools/debug"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GameController(c echo.Context) error {
	debug.Log("front<GameController>")

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

	if student.ID == 0 && !student.IsLicenced() {
		return c.Redirect(http.StatusFound, "/cms/don/")
	} else {
		var step models.Step
		stepID, _ := strconv.Atoi(c.QueryParam("id"))
		stepOrder, _ := strconv.Atoi(c.QueryParam("order"))

		if stepID != 0 {
			db.Set("gorm:auto_preload", true).Where("ID = ?", stepID).First(&step)
		} else  {
			db.Set("gorm:auto_preload", true).Where("order = ?", stepOrder).First(&step)
		}

		if stepID == 0 && stepOrder == 0 {
			var steps []models.Step
			db.Set("gorm:auto_preload", true).Order("order").Find(&steps)


			var notes []models.Note
			db.Where("student_refer = ?", student.ID).Find(&notes)

			for i := range notes {
				for j := range steps {
					if steps[j].ID == notes[i].StepRefer {
						steps[j].Note = notes[i].Note
					}
				}
			}

			var stepsType []models.Step
			db.Select("DISTINCT(type)").Find(&stepsType)

			return c.Render(http.StatusOK, site.Host+"abacusfingers.local<game>", map[string]interface{}{
				"title":      site.Title,
				"steps":      steps,
				"steps_type": stepsType,
				"student":    student,
			})
		} else {
			//TODO: ajouter && student.IsLicenced
			return c.Render(http.StatusOK, site.Host+"abacusfingers.local<_"+step.Type+">", map[string]interface{}{
				"title":   site.Title,
				"step":    step,
				"student": student,
			})
		}
	}
}

func PGameController(c echo.Context) error {
	debug.Log("front<PGameController>")

	db := models.GetDataBase()

	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	student := models.StudentSession(db, c)

	if student.ID == 0 && !student.IsLicenced() {
		return c.Redirect(http.StatusFound, "/cms/don/")
	} else {
		var step models.Step
		db.Where("ID = ?", c.FormValue("gameID")).First(&step)

		errorCounter, _ := strconv.Atoi(c.FormValue("errorCounter"))

		var note models.Note
		db.Where("student_refer = ? AND step_refer = ?", student.ID, step.ID).First(&note)

		note.StepRefer = step.ID
		note.StudentRefer = student.ID

		if (3 - errorCounter) >= note.Note {
			note.Note = 3 - errorCounter
		}

		db.Save(&note)
	}

	return nil
}