package server

import (
	"net/http"
	"time"

	"github.com/AxlLove/go_final_project/pkg/api"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	HTTP *http.Server
}

func CreateServer(addr string, static string) *Server {
	r := chi.NewRouter()

	api.Init(r)
	r.Handle("/*", http.FileServer(http.Dir(static)))

	s := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{HTTP: s}
}
