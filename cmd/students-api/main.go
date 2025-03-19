package main

import (
	"log"
	"net/http"

	"github.com/amreshpro/students-api/internal/config"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("welcome to students api"))
	})
	// setup server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	println("Server started")
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal("failed to start server")
	}
}
