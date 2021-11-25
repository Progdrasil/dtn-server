package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgresql://postgres:postgres@localhost:5432/postgres"
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	log.Println("Connected to " + dbUrl)

	defer db.Close()
	server := newServer(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/bundle/{id:[a-z]+}", server.getBundles)
	r.Post("/bundle/{id:[a-z]+}", server.createBundle)
	http.ListenAndServe(":"+port, r)
}

type Server struct {
	db *sql.DB
}

func newServer(db *sql.DB) *Server {
	return &Server{
		db,
	}
}

func (s *Server) getBundles(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) createBundle(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")

}
