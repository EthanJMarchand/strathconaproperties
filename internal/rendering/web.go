package rendering

import (
	"bytes"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
)

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
	tpl := template.New(patterns[0])
	tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField is not implemented")
		},
	})

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	return &Template{tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t *Template) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := t.htmlTpl.Clone()
	if err != nil {
		fmt.Printf("Error cloning template: %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	tmpl = tmpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})
	buf := bytes.NewBuffer(nil)
	err = tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	_, err = io.Copy(w, buf)
	if err != nil {
		fmt.Printf("Error writing response: %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
