package web

import (
	"fmt"
	"github.com/ethanjmarchand/StrathconaProperties/internal/sp"
	"net/http"
)

// HomeHandler is the struct with the ServeHTTP method on it to pass into mux.Handle.
type HomeHandler struct {
	RS sp.RenderingService
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		ActiveTab string
	}{
		ActiveTab: "home",
	}
	ns := []string{"home.gohtml", "layout.gohtml"}
	err := h.RS.Renderer.Render(w, ns, data)
	if err != nil {
		_ = fmt.Errorf("error rendering home page: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewHomeHandler(rs sp.RenderingService) HomeHandler {
	return HomeHandler{
		RS: rs,
	}
}

// ActivesHandler is the struct with the ServeHTTP method on it to pass into mux.Handle.
type ActivesHandler struct {
	RS sp.RenderingService
}

func (h *ActivesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	actives := []sp.Listing{
		{
			ID:    1,
			Beds:  4,
			Baths: 4,
		},
		{
			ID:    2,
			Beds:  4,
			Baths: 4,
		},
		{
			ID:    3,
			Beds:  4,
			Baths: 4,
		},
	}

	data := struct {
		Actives   []sp.Listing
		ActiveTab string
	}{
		Actives:   actives,
		ActiveTab: "actives",
	}
	ns := []string{"actives.gohtml", "layout.gohtml"}
	err := h.RS.Renderer.Render(w, ns, data)
	if err != nil {
		_ = fmt.Errorf("error rendering actives page: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewActivesHandler(rs sp.RenderingService) ActivesHandler {
	return ActivesHandler{
		RS: rs,
	}
}

// ContactHandler is the struct with the ServeHTTP method on it to pass into mux.Handle.
type ContactHandler struct {
	RS sp.RenderingService
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		ActiveTab string
	}{
		ActiveTab: "contact",
	}
	ns := []string{"contact.gohtml", "layout.gohtml"}
	err := h.RS.Renderer.Render(w, ns, data)
	if err != nil {
		_ = fmt.Errorf("error rendering contact page: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewContactHandler(rs sp.RenderingService) ContactHandler {
	return ContactHandler{
		RS: rs,
	}
}

// SignUpHandler is the struct with the ServeHTTP method on it to pass into mux.Handle.
type SignUpHandler struct {
	RS sp.RenderingService
}

func (h *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ns := []string{"signup.gohtml", "layout.gohtml"}
	err := h.RS.Renderer.Render(w, ns, nil)
	if err != nil {
		_ = fmt.Errorf("error rendering signup page: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func NewSignUpHandler(rs sp.RenderingService) SignUpHandler {
	return SignUpHandler{
		RS: rs,
	}
}
