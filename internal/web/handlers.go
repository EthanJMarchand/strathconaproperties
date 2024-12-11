package web

import (
	"github.com/ethanjmarchand/StrathconaProperties/internal/views"
	"net/http"
)

type HomeHandler struct {
	//US   *sp.UserService
	//LS   *sp.ListingService
	TMPL         *views.Template
	TMPLNotFound *views.Template
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		h.TMPLNotFound.Execute(w, nil)
		return
	}
	h.TMPL.Execute(w, nil)
}

func NewHomeHandler(tmpl *views.Template, tmplNotFound *views.Template) HomeHandler {
	return HomeHandler{
		//US:   us,
		//LS:   ls,
		TMPL:         tmpl,
		TMPLNotFound: tmplNotFound,
	}
}

type ActivesHandler struct {
	TMPL *views.Template
}

func (h *ActivesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	actives := []struct {
		Address string
		Beds    int
		Baths   int
		Price   float32
	}{
		{
			Address: "123 Fake street",
			Beds:    3,
			Baths:   4,
			Price:   399900.00,
		},
		{
			Address: "589 Hunters Grove",
			Beds:    3,
			Baths:   4,
			Price:   750000.99,
		},
		{
			Address: "2408 State St",
			Beds:    5,
			Baths:   4,
			Price:   110000.55,
		}, {
			Address: "Test Address 2",
			Beds:    5,
			Baths:   4,
			Price:   110000.55,
		},
	}
	h.TMPL.Execute(w, actives)
}

func NewActivesHandler(tmpl *views.Template) ActivesHandler {
	return ActivesHandler{
		TMPL: tmpl,
	}
}

type ContactHandler struct {
	TMPL *views.Template
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.TMPL.Execute(w, nil)
}

func NewContactHandler(tmpl *views.Template) ContactHandler {
	return ContactHandler{
		TMPL: tmpl,
	}
}
