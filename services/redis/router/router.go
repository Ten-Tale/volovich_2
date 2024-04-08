package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

func NewRouter(rdb *redis.Client) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/students", getStudents(rdb))

	return r
}
