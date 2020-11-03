package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/priyajithc/cart-svc/handler"
	"github.com/rs/cors"
)

func main() {
	log.Println("Starting server at 8081..")
	cHandler := handler.CartHandler()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cart/{id}", cHandler.GetCart).Methods("GET")
	router.HandleFunc("/cart/{id}", cHandler.AddToCart).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
