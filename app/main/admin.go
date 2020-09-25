package main

import (
	"../models"
	"../tools/debug"
	"github.com/joho/godotenv"
	"github.com/qor/admin"
	"net/http"
	"os"
)

func main() {
	_ = godotenv.Load(".env")

	models.Migrate()
	db := models.GetDataBase()
	defer db.Close()
	debug.Log("<%v : db instance close\n", db)

	Admin := admin.New(&admin.AdminConfig{DB: db})

	//TIPS: Put all struct to manage here.
	Admin.AddResource(&models.Site{})
	Admin.AddResource(&models.Student{})
	Admin.AddResource(&models.Lesson{})
	Admin.AddResource(&models.Train{})
	Admin.AddResource(&models.Step{})
	Admin.AddResource(&models.Note{})
	Admin.AddResource(&models.Cms{})
	Admin.AddResource(&models.Order{})

	mux := http.NewServeMux()
	Admin.MountTo("/", mux)
	_ = http.ListenAndServe(":"+os.Getenv("ADMIN_PORT"), mux)
}
