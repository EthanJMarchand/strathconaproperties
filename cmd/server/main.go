package main

import (
	"fmt"
	"github.com/ethanjmarchand/StrathconaProperties/internal/middleware"
	"github.com/ethanjmarchand/StrathconaProperties/internal/templates"
	"github.com/ethanjmarchand/StrathconaProperties/internal/views"
	"github.com/ethanjmarchand/StrathconaProperties/internal/web"
	"log"
	"net/http"
)

func main() {
	// Connect to database.

	// Parse our templates.
	homeTmpl := views.Must(views.ParseFS(templates.FS, "layout.gohtml", "home.gohtml"))
	notFoundTmpl := views.Must(views.ParseFS(templates.FS, "layout.gohtml", "notfound.gohtml"))
	activesTmpl := views.Must(views.ParseFS(templates.FS, "layout.gohtml", "actives.gohtml"))
	contactTmpl := views.Must(views.ParseFS(templates.FS, "layout.gohtml", "contact.gohtml"))

	// Create services

	// Register handlers and pass in services
	homeHandler := web.NewHomeHandler(homeTmpl, notFoundTmpl)
	activesHandler := web.NewActivesHandler(activesTmpl)
	contactHandler := web.NewContactHandler(contactTmpl)
	// Register routes
	mux := http.NewServeMux()
	mux.Handle("GET /", middleware.Logger(&homeHandler))
	mux.Handle("GET /actives", &activesHandler)
	mux.Handle("GET /contact", &contactHandler)

	// Serve the static assets
	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))
	// Start the HTTP server
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
