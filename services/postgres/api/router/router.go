package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func NewRouter(_ *pgx.Conn) *chi.Mux {
	r := chi.NewRouter()

	return r
}
