package server

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	db   *sql.DB
	port string
}

func New(db *sql.DB, port string) *Server {
	return &Server{
		db:   db,
		port: port,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	// routes
	s.registerRoutes(mux)

	addr := ":" + s.port

	log.Println("Server started on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", s.healthHandler)

	// сюда позже добавим:
	// mux.HandleFunc("/transfer", s.transferHandler)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
