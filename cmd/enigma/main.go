package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/undeadops/enigma/pkg/config"
	"github.com/undeadops/enigma/pkg/db"
	"github.com/undeadops/enigma/pkg/questions"

	"github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"
)

// Debug - Enable debug logging
var Debug = flag.Bool("debug", false, "Enable Debug Logging")

func router() *chi.Mux {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5, ""),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
		cors.Handler,
	)

	return router
}

func main() {
	flag.Parse()

	c := config.New()

	port := os.Getenv("PORT")
	honeyKey := os.Getenv("HONEY_KEY")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Setup Beeline
	beeline.Init(beeline.Config{
		WriteKey:    honeyKey,
		Dataset:     "myData",
		ServiceName: "enigma",
		// ... additional optional configuration ...
	})

	// ensure everything gets sent off before we exit
	defer beeline.Close()

	// Setup Connection timeouts
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Connect to database
	fmt.Printf("Connecting to Database\n")
	qrepo, err := db.NewQuestionsRepo(ctx, c.URI, c.DB)
	if err != nil {
		// Implement better health checking/retry here or in lib
		log.Fatalf("Cannot set up Database: %v", err)
	}

	qhandler := questions.NewHandler(qrepo)

	// Setup Router
	router := router()

	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", hnynethttp.WrapHandlerFunc(index))

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/questions", questions.Router(qhandler))
	})

	fmt.Printf("Starting up Webserver\n")
	log.Fatal(http.ListenAndServe(":"+port, hnynethttp.WrapHandler(router)))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
