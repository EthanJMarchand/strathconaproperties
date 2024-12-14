package rendering

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

func (t *Template) Render(w http.ResponseWriter, names []string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	for _, name := range names {
		if err := t.htmlTpl.ExecuteTemplate(w, name, data); err != nil {
			return err
		}
	}
	return nil
}
