package front

import (
	"../../models"
	"../../tools/debug"
	"github.com/labstack/echo"
)

var _sites []models.Site

func Front(e *echo.Echo) {
	if len(_sites) == 0 {
		db := models.GetDataBase()
		defer db.Close()
		debug.Log("<%v : db instance close\n", db)
		db.Find(&_sites)
	}

	e.GET("/", HomeController)

	e.GET("/cms/*", CmsController)

	e.GET("/nous-contacter", ContactController)
	e.GET("/nous-contacter/", ContactController)
	e.POST("/sendmail", PContactController)

	e.GET("/jeu", GameController)
	e.GET("/jeu/", GameController)
	e.POST("/jeu", PGameController)
	e.POST("/jeu/", PGameController)

	e.GET("/connexion", LoginController)
	e.GET("/connexion/", LoginController)
	e.POST("/connexion", PLoginController)
	e.POST("/connexion/", PLoginController)
	e.GET("/deconnexion", DisconnectionController)
	e.GET("/deconnexion/", DisconnectionController)

	e.GET("/creation-compte", NewAccountController)
	e.GET("/creation-compte/", NewAccountController)
	e.POST("/creation-compte", PNewAccountController)
	e.POST("/creation-compte/", PNewAccountController)
	e.GET("/mon-compte", AccountController)
	e.GET("/mon-compte/", AccountController)

	e.GET("/commander", OrderController)
	e.GET("/commander/", OrderController)
	e.POST("/commander", POrderController)
	e.POST("/commander/", POrderController)



   e.Static("valjeu", "../../public/assets/validation/")
 
}
