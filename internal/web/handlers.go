package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethanjmarchand/StrathconaProperties/internal/sp"
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
	h.RS.Renderer.Render(w, r, data)
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
		{
			ID:    4,
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
	h.RS.Renderer.Render(w, r, data)
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
	h.RS.Renderer.Render(w, r, data)
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
	email := r.FormValue("email")
	data := struct {
		Email     string
		ActiveTab string
	}{
		Email:     email,
		ActiveTab: "signup",
	}
	h.RS.Renderer.Render(w, r, data)
}

func NewSignUpHandler(rs sp.RenderingService) SignUpHandler {
	return SignUpHandler{
		RS: rs,
	}
}

type ProcessSignUpHandler struct {
	US sp.UserService
}

func (h *ProcessSignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	phone := r.PostFormValue("phone")
	password := r.PostFormValue("password")
	if firstName == "" || lastName == "" || email == "" || phone == "" {
		http.Error(w, "First name and last name and phone are required", http.StatusBadRequest)
		return
	}
	user, err := h.US.DB.Create(r.Context(), firstName, lastName, email, phone, password)
	if err != nil {
		fmt.Printf("Error creating user: %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Created user: %v", user)

}

func NewProcessSignUpHandler(us sp.UserService) ProcessSignUpHandler {
	return ProcessSignUpHandler{
		US: us,
	}
}

type SignInHandler struct {
	RS sp.RenderingService
}

func (h *SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		ActiveTab string
		email     string
	}{
		ActiveTab: "signin",
	}
	data.email = r.FormValue("email")
	h.RS.Renderer.Render(w, r, data)
}

func NewSignInHandler(rs sp.RenderingService) SignInHandler {
	return SignInHandler{
		RS: rs,
	}
}

type ProcessSignInHandler struct {
	US sp.UserService
}

func (h *ProcessSignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := h.US.DB.Authenticate(r.Context(), email, password)
	if err != nil {
		log.Printf("Error authenticating user: %v", err)
		// TODO: return an error to the user with proper context and info.
		http.Error(w, "failed authentication", http.StatusUnauthorized)
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Authenticated user: %v", user)
}

func NewProcessSignInHandler(us sp.UserService) ProcessSignInHandler {
	return ProcessSignInHandler{
		US: us,
	}
}
