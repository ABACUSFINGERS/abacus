package front

import (
	"../../models"
	"../../tools/debug"
	"github.com/labstack/echo"
	"net/http"
	"log"
	"net/smtp"
	"os"
)

func ContactController(c echo.Context) error {
	debug.Log("front<ContactController>")

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

	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<contact>", map[string]interface{}{
		"title":   site.Title,
		"student": student,
		"host": os.Getenv("SITE_HOST"),
	})
}


func PContactController(c echo.Context) error {  

	debug.Log("front<PContactController>")
	
	name      := c.FormValue("name")
	email     := c.FormValue("email")
	subject   := c.FormValue("subject")
	comment   := c.FormValue("comment")

	var site models.Site


	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "abacusfingers@gmail.com",  os.Getenv("PASSWORD_SMTP"), "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"abacusfingers@gmail.com"}

	msg := []byte("To: "+"abacusfingers@gmail.com"+"\r\n" +
		"Subject: "+subject+"\r\n" +
		"\r\n" +
		"name= "+name+" email=  "+email+"comment ="+comment+ "\r\n")

	err := smtp.SendMail("smtp.gmail.com:465", auth, "abacusfingers@gmail.com", to, msg)

	if err != nil {
		log.Fatal(err)
		return c.Render(http.StatusOK, site.Host+"abacusfingers.local<contact>", map[string]interface{}{
			"message":  false,
		});
	}
	
	return c.Render(http.StatusOK, site.Host+"abacusfingers.local<contact>", map[string]interface{}{
		"message":   true,
		
	});
}
