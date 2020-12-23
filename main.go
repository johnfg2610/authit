package main

import (
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
)

var (
	serverOptions = &server.Options{
		//TODO customize logger maybe using zerolog?
		RequestLogger: requestlog.NewNCSALogger(os.Stdout, func(error) {}),
	}
)

func main() {
	router := chi.NewRouter()
	srv := server.New(router, serverOptions)

	router.Use(middleware.RequestID)
	//router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}
