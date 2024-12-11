package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	return &Template{tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		_ = fmt.Errorf("error executing temmplate: %w", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
