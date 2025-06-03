package api

import (
	"fmt"
	"net/http"

	"ru.axel.pro.project/config"
)

func Api_v1(server *http.ServeMux, app *config.Application) {
	server.HandleFunc("GET /api/v1/get/menu/side", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"test\":\"test\"}")
	})
}
