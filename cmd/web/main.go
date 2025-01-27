package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mauFade/bookings-u/pkg/config"
	"github.com/mauFade/bookings-u/pkg/handlers"
	"github.com/mauFade/bookings-u/pkg/renders"
)

const PORT = ":9091"

func main() {
	var app config.AppConfig

	tc, err := renders.CreateTemplate()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
