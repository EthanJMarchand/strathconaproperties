package main

import (
	"fmt"
	"github.com/ethanjmarchand/StrathconaProperties/internal/middleware"
	"github.com/ethanjmarchand/StrathconaProperties/internal/rendering"
	"github.com/ethanjmarchand/StrathconaProperties/internal/sp"
	"github.com/ethanjmarchand/StrathconaProperties/internal/templates"
	"github.com/ethanjmarchand/StrathconaProperties/internal/web"
	"log"
	"net/http"
)

func main() {
	// Connect to database.

	// Parse our templates.
	tmpl := rendering.Must(rendering.ParseFS(templates.FS, "*.gohtml"))

	// Create services
	rs := sp.NewRenderingService(tmpl)

	// Register handlers and pass in services and templates
	homeHandler := web.NewHomeHandler(rs)
	activesHandler := web.NewActivesHandler(rs)
	contactHandler := web.NewContactHandler(rs)
	signUpHandler := web.NewSignUpHandler(rs)

	// Register routes
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.Logger(&homeHandler))
	mux.Handle("GET /actives", &activesHandler)
	mux.Handle("GET /contact", &contactHandler)
	mux.Handle("GET /signup", &signUpHandler)

	// Serve the static assets
	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	// Start the HTTP server
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
