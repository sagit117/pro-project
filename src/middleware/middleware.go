package middleware

import (
	"net/http"
	"time"

	"ru.axel.pro.project/config"
)

func Middleware(next http.Handler, app *config.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// cors
		w.Header().Set("Access-Control-Request-Method", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		start := time.Now()
		next.ServeHTTP(w, req)

		Logging(app, w, req, start)

	})
}

func Logging(app *config.Application, w http.ResponseWriter, req *http.Request, start time.Time) {
	app.Logger.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
}
