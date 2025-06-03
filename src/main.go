package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"ru.axel.pro.project/api"
	"ru.axel.pro.project/config"
	"ru.axel.pro.project/middleware"
)

func main() {
	var server_cfg config.ConfigServer

	flag.IntVar(&server_cfg.Port, "p", 3000, "Server port")
	flag.StringVar(&server_cfg.StaticPath, "s", "../static", "Path to static files")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &config.Application{
		ConfigServer: &server_cfg,
		Logger:       logger,
	}

	fs := http.FileServer(http.Dir(server_cfg.StaticPath))
	mux := http.NewServeMux()

	/* Routing */
	mux.Handle("GET /", fs)
	api.Api_v1(mux, app)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", server_cfg.Port),
		Handler:      middleware.Middleware(mux, app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Listening on :%d...", server_cfg.Port)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
