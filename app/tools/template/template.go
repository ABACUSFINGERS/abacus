package template

import (
	"errors"
	"../debug"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Registry struct {
	Templates map[string]*template.Template
}

func (t *Registry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func Process() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	err := filepath.Walk(os.Getenv("TEMPLATES_FOLDER"),
		func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() == false {
				file := strings.TrimSuffix(path, filepath.Ext(path))
				name := filepath.Base(filepath.Dir(path)) + "<" + filepath.Base(file) + ">"

				/*

					segments := strings.Split(strings.TrimSuffix(path, filepath.Ext(path)), "/")[0]
					file := string(segments[0])

				*/

				if filepath.Base(file) != "base" {
					debug.Log(">%v : template file registered with name %v\n", path, name)

					//templates[name] = template.Must(template.ParseFiles(path, filepath.Dir(path)+"/menu_toolbar.gohtml"))

					if string(filepath.Base(file)[0]) == "_" {
						templates[name] = template.Must(template.ParseFiles(path, "../../public/templates/abacusfingers.local/menutoolbar.html"))
					} else {
						switch filepath.Base(file) {
						case "cms":
							templates[name] = template.Must(template.ParseFiles(filepath.Dir(path)+"/cms.html", filepath.Dir(path)+"/base.html"))
							break
						default:
							templates[name] = template.Must(template.ParseFiles(path, filepath.Dir(path)+"/base.html"))
                           
						}
					}

				}
			}
			return nil
		})
	if err != nil {
		debug.Error(">%v : templates process error\n", err)
	}

	return templates
}
