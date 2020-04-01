package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/undeadops/enigma/pkg/config"
	"github.com/undeadops/enigma/pkg/db"
	"github.com/undeadops/enigma/pkg/questions"
)

const defaultPortVariable = "PORT"
const defaultPort = "3000"

// Debug - Enable debug logging
var Debug = flag.Bool("debug", false, "Enable Debug Logging")

func router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	return router
}

func main() {
	flag.Parse()

	c := config.New()

	// Setup Connection timeouts
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Connect to database
	qrepo, err := db.NewQuestionsRepo(ctx, c.URI, c.DB)
	if err != nil {
		// Implement better health checking/retry here or in lib
		log.Fatalf("Cannot set up Database: %v", err)
	}

	qhandler := questions.NewHandler(qrepo)

	// Setup Router
	router := router()
	router.Route("api/v1", func(r chi.Router) {
		r.Mount("/questions", questions.Router(qhandler))
	})

	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
