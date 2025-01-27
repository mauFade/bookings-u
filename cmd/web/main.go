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
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
