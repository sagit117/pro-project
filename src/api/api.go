package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ru.axel.pro.project/config"
)

type MenuSide struct {
	Project string
}

func Api_v1(server *http.ServeMux, app *config.Application) {
	server.HandleFunc("GET /api/v1/get/menu/side", func(w http.ResponseWriter, r *http.Request) {
		menu := &MenuSide{
			Project: "Проекты",
		}

		data, err := json.Marshal(menu)
		if err != nil {
			app.Logger.Fatal(err)
		}

		fmt.Fprint(w, string(data))
	})
}
