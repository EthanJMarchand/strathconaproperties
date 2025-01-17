package main

import (
	"fmt"
	"github.com/ethanjmarchand/StrathconaProperties/internal/middleware"
	"github.com/ethanjmarchand/StrathconaProperties/internal/rendering"
	"github.com/ethanjmarchand/StrathconaProperties/internal/sp"
	"github.com/ethanjmarchand/StrathconaProperties/internal/startup"
	"github.com/ethanjmarchand/StrathconaProperties/internal/store"
	"github.com/ethanjmarchand/StrathconaProperties/internal/templates"
	"github.com/ethanjmarchand/StrathconaProperties/internal/web"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
)

func main() {
	config, err := startup.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	if err = run(config); err != nil {
		log.Fatal(err)
	}
}

func run(s *startup.Config) error {
	// Connect to database.
	cp, err := store.NewPostgresStore(s)
	if err != nil {
		return fmt.Errorf("could not connect to postgres store: %w", err)
	}
	defer cp.DB.Close()

	// Parse our templates.
	homeTmpl := rendering.Must(rendering.ParseFS(templates.FS, "layout.gohtml", "home.gohtml"))
	activesTmpl := rendering.Must(rendering.ParseFS(templates.FS, "layout.gohtml", "actives.gohtml"))
	contactTmpl := rendering.Must(rendering.ParseFS(templates.FS, "layout.gohtml", "contact.gohtml"))
	signinTmpl := rendering.Must(rendering.ParseFS(templates.FS, "layout.gohtml", "signin.gohtml"))
	signUpTmpl := rendering.Must(rendering.ParseFS(templates.FS, "layout.gohtml", "signup.gohtml"))

	// Create renderingService
	homeRS := sp.NewRenderingService(homeTmpl)
	activesRS := sp.NewRenderingService(activesTmpl)
	contactRS := sp.NewRenderingService(contactTmpl)
	signinRS := sp.NewRenderingService(signinTmpl)
	signupRS := sp.NewRenderingService(signUpTmpl)

	// Create userService
	us := sp.NewUserService(cp)
	// Create listingService

	// Register handlers and pass in services and templates
	homeHandler := web.NewHomeHandler(homeRS)
	activesHandler := web.NewActivesHandler(activesRS)
	contactHandler := web.NewContactHandler(contactRS)
	signUpHandler := web.NewSignUpHandler(signupRS)
	processSignUpHandler := web.NewProcessSignUpHandler(us)
	signInHandler := web.NewSignInHandler(signinRS)
	processSignInHandler := web.NewProcessSignInHandler(us)

	// Register routes
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", middleware.Logger(&homeHandler))
	mux.Handle("GET /actives", &activesHandler)
	mux.Handle("GET /contact", &contactHandler)
	mux.Handle("GET /signup", &signUpHandler)
	mux.Handle("POST /signup", &processSignUpHandler)
	mux.Handle("GET /signin", &signInHandler)
	mux.Handle("POST /signin", &processSignInHandler)

	// Serve the static assets
	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	// Start the HTTP server
	fmt.Println("Listening on port 8080")

	CSRF := csrf.Protect([]byte(s.CSRFSecret), csrf.Secure(false))
	log.Fatal(http.ListenAndServe(":8080", CSRF(mux)))
	return nil
}
