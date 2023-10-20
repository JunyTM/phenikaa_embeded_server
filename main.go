package main

import (
	"embedded/controller"
	"embedded/infrastructure"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Printf("Server running at port: %+v\n", infrastructure.GetAppPort())
	server := &http.Server{
		Addr:           ":" + infrastructure.GetAppPort(),
		Handler:        router(),
		ReadTimeout:    6000 * time.Second,
		WriteTimeout:   6000 * time.Second,
		MaxHeaderBytes: 1 << 30,
	}
	log.Fatal(server.ListenAndServe())
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Declear handler
	socket := controller.NewWebSocketController()
	traffic := controller.NewTrafficController()

	r.HandleFunc("/ws", socket.StartConnect)
	r.Route("/api", func(router chi.Router) {
		router.Get("/traffic", traffic.GetTrafficLight)
		router.Put("/traffic", traffic.UpdateTrafficLight)
	})
	return r
}
