package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ru.axel.pro.project/config"
)

type MenuSideItem struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func Api_v1(server *http.ServeMux, app *config.Application) {
	server.HandleFunc("GET /api/v1/get/menu/app", func(w http.ResponseWriter, r *http.Request) {
		menu := []MenuSideItem{}

		project := MenuSideItem{Name: "Проекты", Link: "./projects"}
		actors := MenuSideItem{Name: "Участники", Link: "./actors"}

		menu = append(menu, project)
		menu = append(menu, actors)

		data, err := json.Marshal(menu)
		if err != nil {
			app.Logger.Fatal(err)
		}

		fmt.Fprint(w, string(data))
	})
}
