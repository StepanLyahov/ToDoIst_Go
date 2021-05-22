package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func RunHTTPServer(handler func(router chi.Router) http.Handler) {
	RunHTTPServerOnAddress(":8090", handler)
}

func RunHTTPServerOnAddress(addr string, handler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	addMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", handler(apiRouter))

	log.Println("It is alive!!!!")

	_ = http.ListenAndServe(addr, rootRouter)
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
}
