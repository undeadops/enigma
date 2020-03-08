package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/undeadops/enigma/pkg/config"
	"github.com/undeadops/enigma/pkg/db"
	"github.com/undeadops/enigma/pkg/questions"
	"github.com/undeadops/enigma/pkg/storage"
)

const defaultPortVariable = "PORT"
const defaultPort = "3000"

// Debug - Enable debug logging
var Debug = flag.Bool("debug", false, "Enable Debug Logging")

func routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/questions", questions.Routes())
	})

	router.Routes()
	return router
}

type env struct {
	qr     storage.QuestionsData
	router *chi.Mux
}

func main() {
	flag.Parse()

	c := config.New()

	// Setup Connection timeouts
	//ctx, _ := context.WithTimeout(context.Background(), time.Second()*10)
	ctx := context.Background()

	// Connect to database
	qr, err := db.SetupQuestionsRepo(ctx, c.URI)
	if err != nil {
		// Implement better health checking/retry here or in lib
		log.Fatalf("Cannot set up Database: %v", err)
	}

	s := env{qr: qr, router: routes()}

	//server := &api.Server{DB: db, Port: defaultPort, Ctx: ctx}
	log.Fatal(http.ListenAndServe(":"+defaultPort, s.router))
}
