package main

import (
	"flycode.go/abacusf/app/controllers"
	"flycode.go/abacusf/app/models"
	"flycode.go/abacusf/app/tools/template"
	session "github.com/ipfans/echo-session"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"time"
)

// Thread fil d'execution d'une séquence d'instructions
func main() {
	_ = godotenv.Load(".env")

	e := echo.New()
	models.Migrate()

	e.Use(middleware.Recover())

	e.Server.ReadTimeout = 20 * time.Second
	e.Server.WriteTimeout = 20 * time.Second

	store := session.NewCookieStore([]byte("secret"))
	e.Use(session.Sessions("GSESSION", store))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Static("assets", os.Getenv("ASSETS_FOLDER"))

	templates := template.Process()
	e.Renderer = &template.Registry{
		Templates: templates,
	}

	controllers.Process(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SITE_PORT")))
	//e.Logger.Fatal(e.StartTLS(":443", "../public/certs/*****.crt", "../public/certs/private.key"))
}
