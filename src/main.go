package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "p", 3000, "API server port")
	flag.Parse()

	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// app := &application{
	// 	config: cfg,
	// 	logger: logger,
	// }

	fs := http.FileServer(http.Dir("../static"))
	mux := http.NewServeMux()

	mux.Handle("GET /", fs)
	mux.HandleFunc("GET /api/v1/get/menu/side", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{ \"test\":\"test\"}")
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Listening on :%d...", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
